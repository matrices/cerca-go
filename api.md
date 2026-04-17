# Auth

Response Types:

- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#AuthContextResponse">AuthContextResponse</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#AuthListEnvironmentsResponse">AuthListEnvironmentsResponse</a>

Methods:

- <code title="get /auth/context">client.Auth.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#AuthService.Context">Context</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#AuthContextResponse">AuthContextResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /auth/environments">client.Auth.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#AuthService.ListEnvironments">ListEnvironments</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#AuthListEnvironmentsParams">AuthListEnvironmentsParams</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#AuthListEnvironmentsResponse">AuthListEnvironmentsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# OAuth

Response Types:

- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#OAuthConnectResponse">OAuthConnectResponse</a>

Methods:

- <code title="post /oauth/connect/{provider}">client.OAuth.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#OAuthService.Connect">Connect</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, provider <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#OAuthConnectParams">OAuthConnectParams</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#OAuthConnectResponse">OAuthConnectResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Credentials

Response Types:

- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#Connection">Connection</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CredentialProvider">CredentialProvider</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CredentialType">CredentialType</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CredentialListResponse">CredentialListResponse</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CredentialDeleteResponse">CredentialDeleteResponse</a>

Methods:

- <code title="get /credentials/{scope}">client.Credentials.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CredentialService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, scope <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CredentialListParams">CredentialListParams</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CredentialListResponse">CredentialListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /credentials/{scope}/{connectionId}">client.Credentials.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CredentialService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, scope <a href="https://pkg.go.dev/builtin#string">string</a>, connectionID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CredentialDeleteResponse">CredentialDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /credentials/{scope}/api-keys">client.Credentials.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CredentialService.NewAPIKey">NewAPIKey</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, scope <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CredentialNewAPIKeyParams">CredentialNewAPIKeyParams</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#Connection">Connection</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Environments

Response Types:

- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#Environment">Environment</a>

## ToolSources

Params Types:

- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#HTTPToolDefinitionParam">HTTPToolDefinitionParam</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#HTTPToolExecutionPolicyParam">HTTPToolExecutionPolicyParam</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#McpToolExecutionPolicyParam">McpToolExecutionPolicyParam</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#ToolApprovalMode">ToolApprovalMode</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#ToolSourceAuthParam">ToolSourceAuthParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#HTTPToolDefinition">HTTPToolDefinition</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#HTTPToolExecutionPolicy">HTTPToolExecutionPolicy</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#HTTPToolSource">HTTPToolSource</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#McpToolExecutionPolicy">McpToolExecutionPolicy</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#McpToolSource">McpToolSource</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#ToolApprovalMode">ToolApprovalMode</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#ToolSource">ToolSource</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#ToolSourceAuth">ToolSourceAuth</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#EnvironmentToolSourceListResponse">EnvironmentToolSourceListResponse</a>

Methods:

- <code title="post /environments/{environmentId}/tool-sources">client.Environments.ToolSources.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#EnvironmentToolSourceService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, environmentID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#EnvironmentToolSourceNewParams">EnvironmentToolSourceNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#ToolSource">ToolSource</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /environments/{environmentId}/tool-sources/{sourceId}">client.Environments.ToolSources.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#EnvironmentToolSourceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, environmentID <a href="https://pkg.go.dev/builtin#string">string</a>, sourceID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#ToolSource">ToolSource</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /environments/{environmentId}/tool-sources/{sourceId}">client.Environments.ToolSources.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#EnvironmentToolSourceService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, environmentID <a href="https://pkg.go.dev/builtin#string">string</a>, sourceID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#EnvironmentToolSourceUpdateParams">EnvironmentToolSourceUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#ToolSource">ToolSource</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /environments/{environmentId}/tool-sources">client.Environments.ToolSources.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#EnvironmentToolSourceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, environmentID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#EnvironmentToolSourceListParams">EnvironmentToolSourceListParams</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#EnvironmentToolSourceListResponse">EnvironmentToolSourceListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /environments/{environmentId}/tool-sources/{sourceId}">client.Environments.ToolSources.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#EnvironmentToolSourceService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, environmentID <a href="https://pkg.go.dev/builtin#string">string</a>, sourceID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Webhooks

Params Types:

- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#WebhookSubscriptionEventType">WebhookSubscriptionEventType</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#WebhookEventType">WebhookEventType</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#WebhookSubscription">WebhookSubscription</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#WebhookSubscriptionCreated">WebhookSubscriptionCreated</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#EnvironmentWebhookListResponse">EnvironmentWebhookListResponse</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#EnvironmentWebhookTestResponse">EnvironmentWebhookTestResponse</a>

Methods:

- <code title="post /environments/{environmentId}/webhooks">client.Environments.Webhooks.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#EnvironmentWebhookService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, environmentID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#EnvironmentWebhookNewParams">EnvironmentWebhookNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#WebhookSubscriptionCreated">WebhookSubscriptionCreated</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /environments/{environmentId}/webhooks/{webhookId}">client.Environments.Webhooks.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#EnvironmentWebhookService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, environmentID <a href="https://pkg.go.dev/builtin#string">string</a>, webhookID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#WebhookSubscription">WebhookSubscription</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /environments/{environmentId}/webhooks/{webhookId}">client.Environments.Webhooks.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#EnvironmentWebhookService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, environmentID <a href="https://pkg.go.dev/builtin#string">string</a>, webhookID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#EnvironmentWebhookUpdateParams">EnvironmentWebhookUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#WebhookSubscription">WebhookSubscription</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /environments/{environmentId}/webhooks">client.Environments.Webhooks.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#EnvironmentWebhookService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, environmentID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#EnvironmentWebhookListParams">EnvironmentWebhookListParams</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#EnvironmentWebhookListResponse">EnvironmentWebhookListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /environments/{environmentId}/webhooks/{webhookId}">client.Environments.Webhooks.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#EnvironmentWebhookService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, environmentID <a href="https://pkg.go.dev/builtin#string">string</a>, webhookID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /environments/{environmentId}/webhooks/{webhookId}/rotate">client.Environments.Webhooks.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#EnvironmentWebhookService.Rotate">Rotate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, environmentID <a href="https://pkg.go.dev/builtin#string">string</a>, webhookID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#WebhookSubscriptionCreated">WebhookSubscriptionCreated</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /environments/{environmentId}/webhooks/{webhookId}/test">client.Environments.Webhooks.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#EnvironmentWebhookService.Test">Test</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, environmentID <a href="https://pkg.go.dev/builtin#string">string</a>, webhookID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#EnvironmentWebhookTestResponse">EnvironmentWebhookTestResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Events

Response Types:

- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#RuntimeWebhookEvent">RuntimeWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#RuntimeWebhookEventType">RuntimeWebhookEventType</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#SubscriptionEvent">SubscriptionEvent</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#EnvironmentEventListResponse">EnvironmentEventListResponse</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#EnvironmentEventSubscribeResponse">EnvironmentEventSubscribeResponse</a>

Methods:

- <code title="get /environments/{envId}/events">client.Environments.Events.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#EnvironmentEventService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, envID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#EnvironmentEventListParams">EnvironmentEventListParams</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#EnvironmentEventListResponse">EnvironmentEventListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /environments/{envId}/events/subscribe">client.Environments.Events.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#EnvironmentEventService.Subscribe">Subscribe</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, envID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#EnvironmentEventSubscribeResponse">EnvironmentEventSubscribeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Cells

Params Types:

- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellConfigurationParam">CellConfigurationParam</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellMetadataParam">CellMetadataParam</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#RuntimeToolApprovalOverrideMode">RuntimeToolApprovalOverrideMode</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#Cell">Cell</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellConfiguration">CellConfiguration</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellMetadata">CellMetadata</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellSummary">CellSummary</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#ExecutionPrincipal">ExecutionPrincipal</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#RuntimeToolApprovalOverrideMode">RuntimeToolApprovalOverrideMode</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellListResponse">CellListResponse</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellDeleteResponse">CellDeleteResponse</a>

Methods:

- <code title="post /cells">client.Cells.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellNewParams">CellNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#Cell">Cell</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /cells/{cellId}">client.Cells.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#Cell">Cell</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /cells/{cellId}">client.Cells.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellUpdateParams">CellUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#Cell">Cell</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /cells">client.Cells.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellListParams">CellListParams</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellListResponse">CellListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /cells/{cellId}">client.Cells.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellDeleteResponse">CellDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Config

Methods:

- <code title="get /cells/{cellId}/config">client.Cells.Config.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellConfigService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellConfiguration">CellConfiguration</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Metadata

Methods:

- <code title="patch /cells/{cellId}/metadata">client.Cells.Metadata.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellMetadataService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellMetadataUpdateParams">CellMetadataUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#Cell">Cell</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Tools

Params Types:

- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#RuntimeToolName">RuntimeToolName</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#RuntimeDiscoveredTool">RuntimeDiscoveredTool</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#RuntimeSourceWarning">RuntimeSourceWarning</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#RuntimeToolDescriptor">RuntimeToolDescriptor</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#RuntimeToolName">RuntimeToolName</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellToolListResponse">CellToolListResponse</a>

Methods:

- <code title="get /cells/{cellId}/tools">client.Cells.Tools.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellToolService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellToolListResponse">CellToolListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## ToolSources

Response Types:

- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellToolSourceListResponse">CellToolSourceListResponse</a>

Methods:

- <code title="get /cells/{cellId}/tool-sources">client.Cells.ToolSources.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellToolSourceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellToolSourceListResponse">CellToolSourceListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Context

Response Types:

- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#ContextEntry">ContextEntry</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#ContextEntrySummary">ContextEntrySummary</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#ContextSearchResult">ContextSearchResult</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellContextListResponse">CellContextListResponse</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellContextSearchResponse">CellContextSearchResponse</a>

Methods:

- <code title="get /cells/{cellId}/context">client.Cells.Context.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellContextService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellContextListParams">CellContextListParams</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellContextListResponse">CellContextListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /cells/{cellId}/context/search">client.Cells.Context.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellContextService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellContextSearchParams">CellContextSearchParams</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellContextSearchResponse">CellContextSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Entries

Response Types:

- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellContextEntryDeleteResponse">CellContextEntryDeleteResponse</a>

Methods:

- <code title="get /cells/{cellId}/context/entry">client.Cells.Context.Entries.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellContextEntryService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellContextEntryGetParams">CellContextEntryGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#ContextEntry">ContextEntry</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /cells/{cellId}/context/entry">client.Cells.Context.Entries.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellContextEntryService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellContextEntryDeleteParams">CellContextEntryDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellContextEntryDeleteResponse">CellContextEntryDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /cells/{cellId}/context/entry">client.Cells.Context.Entries.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellContextEntryService.Put">Put</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellContextEntryPutParams">CellContextEntryPutParams</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#ContextEntry">ContextEntry</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Connections

Params Types:

- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#ToolConnectionMetadataParam">ToolConnectionMetadataParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#AttachedConnection">AttachedConnection</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#ToolConnectionMetadata">ToolConnectionMetadata</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellConnectionListResponse">CellConnectionListResponse</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellConnectionDetachResponse">CellConnectionDetachResponse</a>

Methods:

- <code title="get /cells/{cellId}/connections">client.Cells.Connections.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellConnectionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellConnectionListResponse">CellConnectionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /cells/{cellId}/connections">client.Cells.Connections.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellConnectionService.Attach">Attach</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellConnectionAttachParams">CellConnectionAttachParams</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#AttachedConnection">AttachedConnection</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /cells/{cellId}/connections/{connectionId}">client.Cells.Connections.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellConnectionService.Detach">Detach</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>, connectionID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellConnectionDetachResponse">CellConnectionDetachResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Schedules

Response Types:

- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#ScheduledThread">ScheduledThread</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellScheduleListResponse">CellScheduleListResponse</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellScheduleDeleteResponse">CellScheduleDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellScheduleTriggerResponse">CellScheduleTriggerResponse</a>

Methods:

- <code title="post /cells/{cellId}/schedules">client.Cells.Schedules.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellScheduleService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellScheduleNewParams">CellScheduleNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#ScheduledThread">ScheduledThread</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /cells/{cellId}/schedules/{scheduleId}">client.Cells.Schedules.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellScheduleService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>, scheduleID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellScheduleUpdateParams">CellScheduleUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#ScheduledThread">ScheduledThread</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /cells/{cellId}/schedules">client.Cells.Schedules.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellScheduleService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellScheduleListResponse">CellScheduleListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /cells/{cellId}/schedules/{scheduleId}">client.Cells.Schedules.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellScheduleService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>, scheduleID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellScheduleDeleteResponse">CellScheduleDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /cells/{cellId}/schedules/{scheduleId}/trigger">client.Cells.Schedules.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellScheduleService.Trigger">Trigger</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>, scheduleID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellScheduleTriggerResponse">CellScheduleTriggerResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Threads

Params Types:

- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#ThreadStatus">ThreadStatus</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#ContentBlock">ContentBlock</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#PendingSubThread">PendingSubThread</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#SteerResult">SteerResult</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#Thread">Thread</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#ThreadCompiledContext">ThreadCompiledContext</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#ThreadContextWindow">ThreadContextWindow</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#ThreadMessage">ThreadMessage</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#ThreadStatus">ThreadStatus</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#ThreadSummary">ThreadSummary</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#TokenUsage">TokenUsage</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#Turn">Turn</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellThreadListResponse">CellThreadListResponse</a>

Methods:

- <code title="post /cells/{cellId}/threads">client.Cells.Threads.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellThreadService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellThreadNewParams">CellThreadNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#Thread">Thread</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /cells/{cellId}/threads/{threadId}">client.Cells.Threads.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellThreadService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>, threadID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellThreadGetParams">CellThreadGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#Thread">Thread</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /cells/{cellId}/threads">client.Cells.Threads.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellThreadService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellThreadListParams">CellThreadListParams</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellThreadListResponse">CellThreadListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /cells/{cellId}/threads/{threadId}/cancel">client.Cells.Threads.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellThreadService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>, threadID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#Thread">Thread</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /cells/{cellId}/threads/{threadId}/close">client.Cells.Threads.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellThreadService.Close">Close</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>, threadID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#Thread">Thread</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /cells/{cellId}/threads/{threadId}/compact">client.Cells.Threads.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellThreadService.Compact">Compact</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>, threadID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#Thread">Thread</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /cells/{cellId}/threads/{threadId}/steer">client.Cells.Threads.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellThreadService.Steer">Steer</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>, threadID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellThreadSteerParams">CellThreadSteerParams</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#SteerResult">SteerResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /cells/{cellId}/threads/{threadId}/turns">client.Cells.Threads.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellThreadService.Turn">Turn</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>, threadID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellThreadTurnParams">CellThreadTurnParams</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#Turn">Turn</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Approvals

Methods:

- <code title="post /cells/{cellId}/threads/{threadId}/approvals/{approvalId}">client.Cells.Threads.Approvals.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellThreadApprovalService.Resolve">Resolve</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>, threadID <a href="https://pkg.go.dev/builtin#string">string</a>, approvalID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellThreadApprovalResolveParams">CellThreadApprovalResolveParams</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#ApprovalRequest">ApprovalRequest</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### ApprovalGrants

Response Types:

- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellThreadApprovalGrantDeleteResponse">CellThreadApprovalGrantDeleteResponse</a>

Methods:

- <code title="get /cells/{cellId}/threads/{threadId}/approval-grants">client.Cells.Threads.ApprovalGrants.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellThreadApprovalGrantService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>, threadID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*[]<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#ApprovalGrant">ApprovalGrant</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /cells/{cellId}/threads/{threadId}/approval-grants/{grantId}">client.Cells.Threads.ApprovalGrants.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellThreadApprovalGrantService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>, threadID <a href="https://pkg.go.dev/builtin#string">string</a>, grantID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellThreadApprovalGrantDeleteResponse">CellThreadApprovalGrantDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Logs

Response Types:

- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#ThreadLog">ThreadLog</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellThreadLogListResponse">CellThreadLogListResponse</a>

Methods:

- <code title="get /cells/{cellId}/threads/{threadId}/logs">client.Cells.Threads.Logs.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellThreadLogService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>, threadID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellThreadLogListParams">CellThreadLogListParams</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellThreadLogListResponse">CellThreadLogListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Events

Response Types:

- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellThreadEventListResponse">CellThreadEventListResponse</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellThreadEventSubscribeResponse">CellThreadEventSubscribeResponse</a>

Methods:

- <code title="get /cells/{cellId}/threads/{threadId}/events">client.Cells.Threads.Events.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellThreadEventService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>, threadID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellThreadEventListParams">CellThreadEventListParams</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellThreadEventListResponse">CellThreadEventListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /cells/{cellId}/threads/{threadId}/events/subscribe">client.Cells.Threads.Events.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellThreadEventService.Subscribe">Subscribe</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>, threadID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellThreadEventSubscribeResponse">CellThreadEventSubscribeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Approvals

Response Types:

- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#ApprovalRequest">ApprovalRequest</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellApprovalListResponse">CellApprovalListResponse</a>

Methods:

- <code title="get /cells/{cellId}/approvals">client.Cells.Approvals.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellApprovalService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellApprovalListParams">CellApprovalListParams</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellApprovalListResponse">CellApprovalListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## ApprovalGrants

Response Types:

- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#ApprovalGrant">ApprovalGrant</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellApprovalGrantListResponse">CellApprovalGrantListResponse</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellApprovalGrantDeleteResponse">CellApprovalGrantDeleteResponse</a>

Methods:

- <code title="get /cells/{cellId}/approval-grants">client.Cells.ApprovalGrants.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellApprovalGrantService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellApprovalGrantListParams">CellApprovalGrantListParams</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellApprovalGrantListResponse">CellApprovalGrantListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /cells/{cellId}/approval-grants/{grantId}">client.Cells.ApprovalGrants.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellApprovalGrantService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>, grantID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellApprovalGrantDeleteResponse">CellApprovalGrantDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Logs

Response Types:

- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellLogListResponse">CellLogListResponse</a>

Methods:

- <code title="get /cells/{cellId}/logs">client.Cells.Logs.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellLogService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellLogListParams">CellLogListParams</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellLogListResponse">CellLogListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Sandbox

Response Types:

- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#SandboxExecResult">SandboxExecResult</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#SandboxReadResponse">SandboxReadResponse</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellSandboxWriteResponse">CellSandboxWriteResponse</a>

Methods:

- <code title="post /cells/{cellId}/sandbox/exec">client.Cells.Sandbox.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellSandboxService.Exec">Exec</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellSandboxExecParams">CellSandboxExecParams</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#SandboxExecResult">SandboxExecResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /cells/{cellId}/sandbox/read">client.Cells.Sandbox.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellSandboxService.Read">Read</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellSandboxReadParams">CellSandboxReadParams</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#SandboxReadResponse">SandboxReadResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /cells/{cellId}/sandbox/write">client.Cells.Sandbox.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellSandboxService.Write">Write</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellSandboxWriteParams">CellSandboxWriteParams</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellSandboxWriteResponse">CellSandboxWriteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Events

Response Types:

- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellEventListResponse">CellEventListResponse</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellEventSubscribeResponse">CellEventSubscribeResponse</a>

Methods:

- <code title="get /cells/{cellId}/events">client.Cells.Events.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellEventService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellEventListParams">CellEventListParams</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellEventListResponse">CellEventListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /cells/{cellId}/events/subscribe">client.Cells.Events.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellEventService.Subscribe">Subscribe</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cellID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#CellEventSubscribeResponse">CellEventSubscribeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Models

Response Types:

- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#ModelDescriptor">ModelDescriptor</a>
- <a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#ModelListResponse">ModelListResponse</a>

Methods:

- <code title="get /models">client.Models.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#ModelService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/matrices/cerca-go">cercago</a>.<a href="https://pkg.go.dev/github.com/matrices/cerca-go#ModelListResponse">ModelListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
