package middlewares

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set the trace ID in the request context
		traceID := generateTraceID()
		c.Set("TraceID", traceID)

		// Log the request with trace ID, method, path, and request body
		log.Printf("Received request with TraceID: %s, Method: %s, Path: %s, Request Body: %s", traceID, c.Request.Method, c.Request.URL.Path, getRequestBody(c))

		// Create a response capture writer to get the response body
		responseWriter := NewResponseCaptureWriter(c.Writer)

		// Replace the response writer with the capture writer
		c.Writer = responseWriter

		// Continue processing the request
		c.Next()

		// Log the response with trace ID, response status, and response body
		log.Printf("Sent response with TraceID: %s, Status: %d, Response Body: %s", traceID, c.Writer.Status(), responseWriter.GetResponse())
	}
}

func generateTraceID() string {
	// Generate a unique trace ID for the request using UUID
	traceID := uuid.New().String()

	return traceID
}

func getRequestBody(c *gin.Context) string {
	// Get the request body
	body, _ := ioutil.ReadAll(c.Request.Body)

	// Restore the request body for further processing
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return string(body)
}

// ResponseCaptureWriter captures the response body
type ResponseCaptureWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// NewResponseCaptureWriter creates a new instance of ResponseCaptureWriter
func NewResponseCaptureWriter(w gin.ResponseWriter) *ResponseCaptureWriter {
	return &ResponseCaptureWriter{
		ResponseWriter: w,
		body:           bytes.NewBufferString(""),
	}
}

// Write overrides the Write method to capture the response body
func (w *ResponseCaptureWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// GetResponse returns the captured response body
func (w *ResponseCaptureWriter) GetResponse() string {
	return w.body.String()
}
