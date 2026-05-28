# Batch

Response Types:

- <a href="https://pkg.go.dev/github.com/with-ours/ingest-sdk-go">oursprivacy</a>.<a href="https://pkg.go.dev/github.com/with-ours/ingest-sdk-go#BatchNewResponse">BatchNewResponse</a>

Methods:

- <code title="post /batch">client.Batch.<a href="https://pkg.go.dev/github.com/with-ours/ingest-sdk-go#BatchService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/with-ours/ingest-sdk-go">oursprivacy</a>.<a href="https://pkg.go.dev/github.com/with-ours/ingest-sdk-go#BatchNewParams">BatchNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/with-ours/ingest-sdk-go">oursprivacy</a>.<a href="https://pkg.go.dev/github.com/with-ours/ingest-sdk-go#BatchNewResponse">BatchNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Track

Response Types:

- <a href="https://pkg.go.dev/github.com/with-ours/ingest-sdk-go">oursprivacy</a>.<a href="https://pkg.go.dev/github.com/with-ours/ingest-sdk-go#TrackEventResponse">TrackEventResponse</a>

Methods:

- <code title="post /track">client.Track.<a href="https://pkg.go.dev/github.com/with-ours/ingest-sdk-go#TrackService.Event">Event</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/with-ours/ingest-sdk-go">oursprivacy</a>.<a href="https://pkg.go.dev/github.com/with-ours/ingest-sdk-go#TrackEventParams">TrackEventParams</a>) (\*<a href="https://pkg.go.dev/github.com/with-ours/ingest-sdk-go">oursprivacy</a>.<a href="https://pkg.go.dev/github.com/with-ours/ingest-sdk-go#TrackEventResponse">TrackEventResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Visitor

Response Types:

- <a href="https://pkg.go.dev/github.com/with-ours/ingest-sdk-go">oursprivacy</a>.<a href="https://pkg.go.dev/github.com/with-ours/ingest-sdk-go#VisitorUpsertResponse">VisitorUpsertResponse</a>

Methods:

- <code title="post /identify">client.Visitor.<a href="https://pkg.go.dev/github.com/with-ours/ingest-sdk-go#VisitorService.Upsert">Upsert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/with-ours/ingest-sdk-go">oursprivacy</a>.<a href="https://pkg.go.dev/github.com/with-ours/ingest-sdk-go#VisitorUpsertParams">VisitorUpsertParams</a>) (\*<a href="https://pkg.go.dev/github.com/with-ours/ingest-sdk-go">oursprivacy</a>.<a href="https://pkg.go.dev/github.com/with-ours/ingest-sdk-go#VisitorUpsertResponse">VisitorUpsertResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Experiments

Response Types:

- <a href="https://pkg.go.dev/github.com/with-ours/ingest-sdk-go">oursprivacy</a>.<a href="https://pkg.go.dev/github.com/with-ours/ingest-sdk-go#ExperimentAssignmentResponseUnion">ExperimentAssignmentResponseUnion</a>
- <a href="https://pkg.go.dev/github.com/with-ours/ingest-sdk-go">oursprivacy</a>.<a href="https://pkg.go.dev/github.com/with-ours/ingest-sdk-go#ExperimentPersonalizationResponse">ExperimentPersonalizationResponse</a>

Methods:

- <code title="post /experiments/assignments/{experiment_key}">client.Experiments.<a href="https://pkg.go.dev/github.com/with-ours/ingest-sdk-go#ExperimentService.Assignment">Assignment</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, experimentKey <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/with-ours/ingest-sdk-go">oursprivacy</a>.<a href="https://pkg.go.dev/github.com/with-ours/ingest-sdk-go#ExperimentAssignmentParams">ExperimentAssignmentParams</a>) (\*<a href="https://pkg.go.dev/github.com/with-ours/ingest-sdk-go">oursprivacy</a>.<a href="https://pkg.go.dev/github.com/with-ours/ingest-sdk-go#ExperimentAssignmentResponseUnion">ExperimentAssignmentResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /experiments/personalization">client.Experiments.<a href="https://pkg.go.dev/github.com/with-ours/ingest-sdk-go#ExperimentService.Personalization">Personalization</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/with-ours/ingest-sdk-go">oursprivacy</a>.<a href="https://pkg.go.dev/github.com/with-ours/ingest-sdk-go#ExperimentPersonalizationParams">ExperimentPersonalizationParams</a>) (\*<a href="https://pkg.go.dev/github.com/with-ours/ingest-sdk-go">oursprivacy</a>.<a href="https://pkg.go.dev/github.com/with-ours/ingest-sdk-go#ExperimentPersonalizationResponse">ExperimentPersonalizationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
