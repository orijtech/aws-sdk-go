// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisvideoarchivedmedia

import (
	"io"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/request"
	"go.opencensus.io/trace"
)

const opGetMediaForFragmentList = "GetMediaForFragmentList"

// GetMediaForFragmentListRequest generates a "aws/request.Request" representing the
// client's request for the GetMediaForFragmentList operation. The "output" return
// value will be populated with the request's response once the request completes
// successfuly.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See GetMediaForFragmentList for more information on using the GetMediaForFragmentList
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the GetMediaForFragmentListRequest method.
//    req, resp := client.GetMediaForFragmentListRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/kinesis-video-archived-media-2017-09-30/GetMediaForFragmentList
func (c *KinesisVideoArchivedMedia) GetMediaForFragmentListRequest(input *GetMediaForFragmentListInput) (req *request.Request, output *GetMediaForFragmentListOutput) {
	op := &request.Operation{
		Name:       opGetMediaForFragmentList,
		HTTPMethod: "POST",
		HTTPPath:   "/getMediaForFragmentList",
	}

	if input == nil {
		input = &GetMediaForFragmentListInput{}
	}

	output = &GetMediaForFragmentListOutput{}
	req = c.newRequest(op, input, output)
	return
}

// GetMediaForFragmentList API operation for Amazon Kinesis Video Streams Archived Media.
//
// Gets media for a list of fragments (specified by fragment number) from the
// archived data in a Kinesis video stream.
//
// This operation is only available for the AWS SDK for Java. It is not supported
// in AWS SDKs for other languages.
//
// The following limits apply when using the GetMediaForFragmentList API:
//
//    * A client can call GetMediaForFragmentList up to five times per second
//    per stream.
//
//    * Kinesis Video Streams sends media data at a rate of up to 25 megabytes
//    per second (or 200 megabits per second) during a GetMediaForFragmentList
//    session.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon Kinesis Video Streams Archived Media's
// API operation GetMediaForFragmentList for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeResourceNotFoundException "ResourceNotFoundException"
//   Kinesis Video Streams can't find the stream that you specified.
//
//   * ErrCodeInvalidArgumentException "InvalidArgumentException"
//   A specified parameter exceeds its restrictions, is not supported, or can't
//   be used.
//
//   * ErrCodeClientLimitExceededException "ClientLimitExceededException"
//   Kinesis Video Streams has throttled the request because you have exceeded
//   the limit of allowed client calls. Try making the call later.
//
//   * ErrCodeNotAuthorizedException "NotAuthorizedException"
//   Status Code: 403, The caller is not authorized to perform an operation on
//   the given stream, or the token has expired.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/kinesis-video-archived-media-2017-09-30/GetMediaForFragmentList
func (c *KinesisVideoArchivedMedia) GetMediaForFragmentList(input *GetMediaForFragmentListInput) (*GetMediaForFragmentListOutput, error) {
	req, out := c.GetMediaForFragmentListRequest(input)
	return out, req.Send()
}

// GetMediaForFragmentListWithContext is the same as GetMediaForFragmentList with the addition of
// the ability to pass a context and additional request options.
//
// See GetMediaForFragmentList for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KinesisVideoArchivedMedia) GetMediaForFragmentListWithContext(ctx aws.Context, input *GetMediaForFragmentListInput, opts ...request.Option) (*GetMediaForFragmentListOutput, error) {
	ctx, span := trace.StartSpan(ctx, "aws/kinesisvideoarchivedmedia.(*KinesisVideoArchivedMedia).GetMediaForFragmentList")
	defer span.End()

	req, out := c.GetMediaForFragmentListRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListFragments = "ListFragments"

// ListFragmentsRequest generates a "aws/request.Request" representing the
// client's request for the ListFragments operation. The "output" return
// value will be populated with the request's response once the request completes
// successfuly.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See ListFragments for more information on using the ListFragments
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the ListFragmentsRequest method.
//    req, resp := client.ListFragmentsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/kinesis-video-archived-media-2017-09-30/ListFragments
func (c *KinesisVideoArchivedMedia) ListFragmentsRequest(input *ListFragmentsInput) (req *request.Request, output *ListFragmentsOutput) {
	op := &request.Operation{
		Name:       opListFragments,
		HTTPMethod: "POST",
		HTTPPath:   "/listFragments",
	}

	if input == nil {
		input = &ListFragmentsInput{}
	}

	output = &ListFragmentsOutput{}
	req = c.newRequest(op, input, output)
	return
}

// ListFragments API operation for Amazon Kinesis Video Streams Archived Media.
//
// Returns a list of Fragment objects from the specified stream and start location
// within the archived data.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon Kinesis Video Streams Archived Media's
// API operation ListFragments for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeResourceNotFoundException "ResourceNotFoundException"
//   Kinesis Video Streams can't find the stream that you specified.
//
//   * ErrCodeInvalidArgumentException "InvalidArgumentException"
//   A specified parameter exceeds its restrictions, is not supported, or can't
//   be used.
//
//   * ErrCodeClientLimitExceededException "ClientLimitExceededException"
//   Kinesis Video Streams has throttled the request because you have exceeded
//   the limit of allowed client calls. Try making the call later.
//
//   * ErrCodeNotAuthorizedException "NotAuthorizedException"
//   Status Code: 403, The caller is not authorized to perform an operation on
//   the given stream, or the token has expired.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/kinesis-video-archived-media-2017-09-30/ListFragments
func (c *KinesisVideoArchivedMedia) ListFragments(input *ListFragmentsInput) (*ListFragmentsOutput, error) {
	req, out := c.ListFragmentsRequest(input)
	return out, req.Send()
}

// ListFragmentsWithContext is the same as ListFragments with the addition of
// the ability to pass a context and additional request options.
//
// See ListFragments for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KinesisVideoArchivedMedia) ListFragmentsWithContext(ctx aws.Context, input *ListFragmentsInput, opts ...request.Option) (*ListFragmentsOutput, error) {
	ctx, span := trace.StartSpan(ctx, "aws/kinesisvideoarchivedmedia.(*KinesisVideoArchivedMedia).ListFragments")
	defer span.End()

	req, out := c.ListFragmentsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

// Represents a segment of video or other time-delimited data.
type Fragment struct {
	_ struct{} `type:"structure"`

	// The playback duration or other time value associated with the fragment.
	FragmentLengthInMilliseconds *int64 `type:"long"`

	// The index value of the fragment.
	FragmentNumber *string `min:"1" type:"string"`

	// The total fragment size, including information about the fragment and contained
	// media data.
	FragmentSizeInBytes *int64 `type:"long"`

	// The time stamp from the producer corresponding to the fragment.
	ProducerTimestamp *time.Time `type:"timestamp" timestampFormat:"unix"`

	// The time stamp from the AWS server corresponding to the fragment.
	ServerTimestamp *time.Time `type:"timestamp" timestampFormat:"unix"`
}

// String returns the string representation
func (s Fragment) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Fragment) GoString() string {
	return s.String()
}

// SetFragmentLengthInMilliseconds sets the FragmentLengthInMilliseconds field's value.
func (s *Fragment) SetFragmentLengthInMilliseconds(v int64) *Fragment {
	s.FragmentLengthInMilliseconds = &v
	return s
}

// SetFragmentNumber sets the FragmentNumber field's value.
func (s *Fragment) SetFragmentNumber(v string) *Fragment {
	s.FragmentNumber = &v
	return s
}

// SetFragmentSizeInBytes sets the FragmentSizeInBytes field's value.
func (s *Fragment) SetFragmentSizeInBytes(v int64) *Fragment {
	s.FragmentSizeInBytes = &v
	return s
}

// SetProducerTimestamp sets the ProducerTimestamp field's value.
func (s *Fragment) SetProducerTimestamp(v time.Time) *Fragment {
	s.ProducerTimestamp = &v
	return s
}

// SetServerTimestamp sets the ServerTimestamp field's value.
func (s *Fragment) SetServerTimestamp(v time.Time) *Fragment {
	s.ServerTimestamp = &v
	return s
}

// Describes the time stamp range and time stamp origin of a range of fragments.
type FragmentSelector struct {
	_ struct{} `type:"structure"`

	// The origin of the time stamps to use (Server or Producer).
	//
	// FragmentSelectorType is a required field
	FragmentSelectorType *string `type:"string" required:"true" enum:"FragmentSelectorType"`

	// The range of time stamps to return.
	//
	// TimestampRange is a required field
	TimestampRange *TimestampRange `type:"structure" required:"true"`
}

// String returns the string representation
func (s FragmentSelector) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s FragmentSelector) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *FragmentSelector) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "FragmentSelector"}
	if s.FragmentSelectorType == nil {
		invalidParams.Add(request.NewErrParamRequired("FragmentSelectorType"))
	}
	if s.TimestampRange == nil {
		invalidParams.Add(request.NewErrParamRequired("TimestampRange"))
	}
	if s.TimestampRange != nil {
		if err := s.TimestampRange.Validate(); err != nil {
			invalidParams.AddNested("TimestampRange", err.(request.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetFragmentSelectorType sets the FragmentSelectorType field's value.
func (s *FragmentSelector) SetFragmentSelectorType(v string) *FragmentSelector {
	s.FragmentSelectorType = &v
	return s
}

// SetTimestampRange sets the TimestampRange field's value.
func (s *FragmentSelector) SetTimestampRange(v *TimestampRange) *FragmentSelector {
	s.TimestampRange = v
	return s
}

type GetMediaForFragmentListInput struct {
	_ struct{} `type:"structure"`

	// A list of the numbers of fragments for which to retrieve media. You retrieve
	// these values with ListFragments.
	//
	// Fragments is a required field
	Fragments []*string `type:"list" required:"true"`

	// The name of the stream from which to retrieve fragment media.
	//
	// StreamName is a required field
	StreamName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetMediaForFragmentListInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetMediaForFragmentListInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetMediaForFragmentListInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetMediaForFragmentListInput"}
	if s.Fragments == nil {
		invalidParams.Add(request.NewErrParamRequired("Fragments"))
	}
	if s.StreamName == nil {
		invalidParams.Add(request.NewErrParamRequired("StreamName"))
	}
	if s.StreamName != nil && len(*s.StreamName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("StreamName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetFragments sets the Fragments field's value.
func (s *GetMediaForFragmentListInput) SetFragments(v []*string) *GetMediaForFragmentListInput {
	s.Fragments = v
	return s
}

// SetStreamName sets the StreamName field's value.
func (s *GetMediaForFragmentListInput) SetStreamName(v string) *GetMediaForFragmentListInput {
	s.StreamName = &v
	return s
}

type GetMediaForFragmentListOutput struct {
	_ struct{} `type:"structure" payload:"Payload"`

	// The content type of the requested media.
	ContentType *string `location:"header" locationName:"Content-Type" min:"1" type:"string"`

	// The payload that Kinesis Video Streams returns is a sequence of chunks from
	// the specified stream. For information about the chunks, see PutMedia (docs.aws.amazon.com/acuity/latest/dg/API_dataplane_PutMedia.html).
	// The chunks that Kinesis Video Streams returns in the GetMediaForFragmentList
	// call also include the following additional Matroska (MKV) tags:
	//
	//    * AWS_KINESISVIDEO_FRAGMENT_NUMBER - Fragment number returned in the chunk.
	//
	//    * AWS_KINESISVIDEO_SERVER_SIDE_TIMESTAMP - Server-side time stamp of the
	//    fragment.
	//
	//    * AWS_KINESISVIDEO_PRODUCER_SIDE_TIMESTAMP - Producer-side time stamp
	//    of the fragment.
	//
	// The following tags will be included if an exception occurs:
	//
	//    * AWS_KINESISVIDEO_FRAGMENT_NUMBER - The number of the fragment that threw
	//    the exception
	//
	//    * AWS_KINESISVIDEO_EXCEPTION_ERROR_CODE - The integer code of the exception
	//
	//    * AWS_KINESISVIDEO_EXCEPTION_MESSAGE - A text description of the exception
	Payload io.ReadCloser `type:"blob"`
}

// String returns the string representation
func (s GetMediaForFragmentListOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetMediaForFragmentListOutput) GoString() string {
	return s.String()
}

// SetContentType sets the ContentType field's value.
func (s *GetMediaForFragmentListOutput) SetContentType(v string) *GetMediaForFragmentListOutput {
	s.ContentType = &v
	return s
}

// SetPayload sets the Payload field's value.
func (s *GetMediaForFragmentListOutput) SetPayload(v io.ReadCloser) *GetMediaForFragmentListOutput {
	s.Payload = v
	return s
}

type ListFragmentsInput struct {
	_ struct{} `type:"structure"`

	// Describes the time stamp range and time stamp origin for the range of fragments
	// to return.
	FragmentSelector *FragmentSelector `type:"structure"`

	// The total number of fragments to return. If the total number of fragments
	// available is more than the value specified in max-results, then a ListFragmentsOutput$NextToken
	// is provided in the output that you can use to resume pagination.
	MaxResults *int64 `min:"1" type:"long"`

	// A token to specify where to start paginating. This is the ListFragmentsOutput$NextToken
	// from a previously truncated response.
	NextToken *string `min:"1" type:"string"`

	// The name of the stream from which to retrieve a fragment list.
	//
	// StreamName is a required field
	StreamName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ListFragmentsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListFragmentsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListFragmentsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListFragmentsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(request.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("NextToken", 1))
	}
	if s.StreamName == nil {
		invalidParams.Add(request.NewErrParamRequired("StreamName"))
	}
	if s.StreamName != nil && len(*s.StreamName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("StreamName", 1))
	}
	if s.FragmentSelector != nil {
		if err := s.FragmentSelector.Validate(); err != nil {
			invalidParams.AddNested("FragmentSelector", err.(request.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetFragmentSelector sets the FragmentSelector field's value.
func (s *ListFragmentsInput) SetFragmentSelector(v *FragmentSelector) *ListFragmentsInput {
	s.FragmentSelector = v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *ListFragmentsInput) SetMaxResults(v int64) *ListFragmentsInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListFragmentsInput) SetNextToken(v string) *ListFragmentsInput {
	s.NextToken = &v
	return s
}

// SetStreamName sets the StreamName field's value.
func (s *ListFragmentsInput) SetStreamName(v string) *ListFragmentsInput {
	s.StreamName = &v
	return s
}

type ListFragmentsOutput struct {
	_ struct{} `type:"structure"`

	// A list of fragment numbers that correspond to the time stamp range provided.
	Fragments []*Fragment `type:"list"`

	// If the returned list is truncated, the operation returns this token to use
	// to retrieve the next page of results. This value is null when there are no
	// more results to return.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListFragmentsOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListFragmentsOutput) GoString() string {
	return s.String()
}

// SetFragments sets the Fragments field's value.
func (s *ListFragmentsOutput) SetFragments(v []*Fragment) *ListFragmentsOutput {
	s.Fragments = v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListFragmentsOutput) SetNextToken(v string) *ListFragmentsOutput {
	s.NextToken = &v
	return s
}

// The range of time stamps for which to return fragments.
type TimestampRange struct {
	_ struct{} `type:"structure"`

	// The ending time stamp in the range of time stamps for which to return fragments.
	//
	// EndTimestamp is a required field
	EndTimestamp *time.Time `type:"timestamp" timestampFormat:"unix" required:"true"`

	// The starting time stamp in the range of time stamps for which to return fragments.
	//
	// StartTimestamp is a required field
	StartTimestamp *time.Time `type:"timestamp" timestampFormat:"unix" required:"true"`
}

// String returns the string representation
func (s TimestampRange) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s TimestampRange) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TimestampRange) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "TimestampRange"}
	if s.EndTimestamp == nil {
		invalidParams.Add(request.NewErrParamRequired("EndTimestamp"))
	}
	if s.StartTimestamp == nil {
		invalidParams.Add(request.NewErrParamRequired("StartTimestamp"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetEndTimestamp sets the EndTimestamp field's value.
func (s *TimestampRange) SetEndTimestamp(v time.Time) *TimestampRange {
	s.EndTimestamp = &v
	return s
}

// SetStartTimestamp sets the StartTimestamp field's value.
func (s *TimestampRange) SetStartTimestamp(v time.Time) *TimestampRange {
	s.StartTimestamp = &v
	return s
}

const (
	// FragmentSelectorTypeProducerTimestamp is a FragmentSelectorType enum value
	FragmentSelectorTypeProducerTimestamp = "PRODUCER_TIMESTAMP"

	// FragmentSelectorTypeServerTimestamp is a FragmentSelectorType enum value
	FragmentSelectorTypeServerTimestamp = "SERVER_TIMESTAMP"
)
