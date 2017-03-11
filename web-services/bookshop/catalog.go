package catalog

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/kavirajk/bookshop/catalog"
)
// DEFINE OMIT

package catalog

type Service interface {
	// Search books based on free text
	Search(ctx context.Context, query string) ([]Book, error) // HL

	// List available items based on limit and offset.
	List(ctx context.Context, order string, limit, offset int) ([]Book, int, error) // HL

	// Get details about single book
	Get(ctx context.Context, id string) (Book, error) // HL
}

// DEFINEEND OMIT

// IMPL OMIT
type basicService struct {
	r Repo
}

// NewCatalogService return basic Service implementation.
func NewService(r Repo) Service { // HL
	return basicService{r}
}

// Search return books that matches with query.
func (s basicService) Search(ctx context.Context, query string) ([]Book, error) {
	return s.r.Search(query)
}

// Get return a book for the matched ID. Empty book incase of non-error.
func (s basicService) Get(ctx context.Context, ID string) (Book, error) {
	return s.r.GetByID(ID)
}

// IMPLEND OMIT

// MID OMIT
// Middleware is a service middleware
type Middleware func(Service) Service

// MIDEND OMIT

// LOG OMIT
type loggingService struct {
	logger log.Logger // HL
	next   Service    // HL
}

func LoggingMiddleware(logger log.Logger) Middleware { // HL
	return func(next Service) Service {
		return loggingService{
			logger: logger,
			next:   next,
		}
	}
}

func (s loggingService) Search(ctx context.Context, query string) (books []Book, err error) {
	defer func(begin time.Time) {
		_ = s.logger.Log(
			"method", "search",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return s.next.Search(ctx, query) // HL
}

// LOGEND OMIT
// ENDP1 OMIT
func makeSearchEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(searchRequest)
		books, err := svc.Search(req.Q)
		if err != nil {
			return searchResponse{books, err.Error()}, nil
		}
		return uppercaseResponse{books, ""}, nil
	}
}

// ENDP1END OMIT
// ENDP2 OMIT
type searchRequest struct {
	Q string `json:"q"`
}

type searchResponse struct {
	Books []Book `json:"books,omitempty"`
	Error error  `json:"error,omitempty"`
}

// ENDP2END OMIT

// TRANS OMIT
var searchHandler = httptransport.NewServer(
	ctx,
	makeSearchEndpoint(svc),
	decodeSearchRequest,
	encodeResponse,
)

func decodeSearchRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	q := req.FormValue("q")
	if strings.TrimSpace(q) == "" {
		return nil, ErrEmptyQuery
	}
	return searchRequest{
		Q: q,
	}, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// TRANSEND OMIT

// MAIN OMIT
func main() {
	// Other dependencies intialization

	ctx := context.Background()

	svc := catalog.NewService(repo)
	svc = loggingMiddleware(logger)(svc)
	svc = instrumentingMiddleware(requestCount, requestLatency, countResult)(svc)

	searchHandler := httptransport.NewServer(
		ctx,
		makeSearchEndpoint(svc),
		decodeSearchRequest,
		encodeResponse,
	)

	http.Handle("/search", searchHandler)
	http.Handle("/metrics", stdprometheus.Handler())
}

// MAINEND OMIT
