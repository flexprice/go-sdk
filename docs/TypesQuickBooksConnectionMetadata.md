# TypesQuickBooksConnectionMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** | Managed internally - set after auth code exchange or token refresh | [optional] 
**AuthCode** | Pointer to **string** | Optional - for initial setup via auth code (will be cleared after token exchange) | [optional] 
**ClientId** | Pointer to **string** | Required for initial connection setup | [optional] 
**ClientSecret** | Pointer to **string** | OAuth Client Secret (encrypted) | [optional] 
**Environment** | Pointer to **string** | \&quot;sandbox\&quot; or \&quot;production\&quot; | [optional] 
**IncomeAccountId** | Pointer to **string** | Optional configuration | [optional] 
**OauthSessionData** | Pointer to **string** | Temporary OAuth session data (only used during OAuth flow, cleared after completion) | [optional] 
**RealmId** | Pointer to **string** | QuickBooks Company ID (not encrypted) | [optional] 
**RedirectUri** | Pointer to **string** | OAuth Redirect URI (temporary) | [optional] 
**RefreshToken** | Pointer to **string** | OAuth Refresh Token (encrypted) | [optional] 
**WebhookVerifierToken** | Pointer to **string** | Webhook security | [optional] 

## Methods

### NewTypesQuickBooksConnectionMetadata

`func NewTypesQuickBooksConnectionMetadata() *TypesQuickBooksConnectionMetadata`

NewTypesQuickBooksConnectionMetadata instantiates a new TypesQuickBooksConnectionMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesQuickBooksConnectionMetadataWithDefaults

`func NewTypesQuickBooksConnectionMetadataWithDefaults() *TypesQuickBooksConnectionMetadata`

NewTypesQuickBooksConnectionMetadataWithDefaults instantiates a new TypesQuickBooksConnectionMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *TypesQuickBooksConnectionMetadata) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *TypesQuickBooksConnectionMetadata) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *TypesQuickBooksConnectionMetadata) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *TypesQuickBooksConnectionMetadata) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetAuthCode

`func (o *TypesQuickBooksConnectionMetadata) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *TypesQuickBooksConnectionMetadata) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *TypesQuickBooksConnectionMetadata) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.

### HasAuthCode

`func (o *TypesQuickBooksConnectionMetadata) HasAuthCode() bool`

HasAuthCode returns a boolean if a field has been set.

### GetClientId

`func (o *TypesQuickBooksConnectionMetadata) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *TypesQuickBooksConnectionMetadata) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *TypesQuickBooksConnectionMetadata) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *TypesQuickBooksConnectionMetadata) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *TypesQuickBooksConnectionMetadata) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *TypesQuickBooksConnectionMetadata) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *TypesQuickBooksConnectionMetadata) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *TypesQuickBooksConnectionMetadata) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetEnvironment

`func (o *TypesQuickBooksConnectionMetadata) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *TypesQuickBooksConnectionMetadata) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *TypesQuickBooksConnectionMetadata) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *TypesQuickBooksConnectionMetadata) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetIncomeAccountId

`func (o *TypesQuickBooksConnectionMetadata) GetIncomeAccountId() string`

GetIncomeAccountId returns the IncomeAccountId field if non-nil, zero value otherwise.

### GetIncomeAccountIdOk

`func (o *TypesQuickBooksConnectionMetadata) GetIncomeAccountIdOk() (*string, bool)`

GetIncomeAccountIdOk returns a tuple with the IncomeAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeAccountId

`func (o *TypesQuickBooksConnectionMetadata) SetIncomeAccountId(v string)`

SetIncomeAccountId sets IncomeAccountId field to given value.

### HasIncomeAccountId

`func (o *TypesQuickBooksConnectionMetadata) HasIncomeAccountId() bool`

HasIncomeAccountId returns a boolean if a field has been set.

### GetOauthSessionData

`func (o *TypesQuickBooksConnectionMetadata) GetOauthSessionData() string`

GetOauthSessionData returns the OauthSessionData field if non-nil, zero value otherwise.

### GetOauthSessionDataOk

`func (o *TypesQuickBooksConnectionMetadata) GetOauthSessionDataOk() (*string, bool)`

GetOauthSessionDataOk returns a tuple with the OauthSessionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthSessionData

`func (o *TypesQuickBooksConnectionMetadata) SetOauthSessionData(v string)`

SetOauthSessionData sets OauthSessionData field to given value.

### HasOauthSessionData

`func (o *TypesQuickBooksConnectionMetadata) HasOauthSessionData() bool`

HasOauthSessionData returns a boolean if a field has been set.

### GetRealmId

`func (o *TypesQuickBooksConnectionMetadata) GetRealmId() string`

GetRealmId returns the RealmId field if non-nil, zero value otherwise.

### GetRealmIdOk

`func (o *TypesQuickBooksConnectionMetadata) GetRealmIdOk() (*string, bool)`

GetRealmIdOk returns a tuple with the RealmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmId

`func (o *TypesQuickBooksConnectionMetadata) SetRealmId(v string)`

SetRealmId sets RealmId field to given value.

### HasRealmId

`func (o *TypesQuickBooksConnectionMetadata) HasRealmId() bool`

HasRealmId returns a boolean if a field has been set.

### GetRedirectUri

`func (o *TypesQuickBooksConnectionMetadata) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *TypesQuickBooksConnectionMetadata) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *TypesQuickBooksConnectionMetadata) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *TypesQuickBooksConnectionMetadata) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.

### GetRefreshToken

`func (o *TypesQuickBooksConnectionMetadata) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *TypesQuickBooksConnectionMetadata) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *TypesQuickBooksConnectionMetadata) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *TypesQuickBooksConnectionMetadata) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetWebhookVerifierToken

`func (o *TypesQuickBooksConnectionMetadata) GetWebhookVerifierToken() string`

GetWebhookVerifierToken returns the WebhookVerifierToken field if non-nil, zero value otherwise.

### GetWebhookVerifierTokenOk

`func (o *TypesQuickBooksConnectionMetadata) GetWebhookVerifierTokenOk() (*string, bool)`

GetWebhookVerifierTokenOk returns a tuple with the WebhookVerifierToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookVerifierToken

`func (o *TypesQuickBooksConnectionMetadata) SetWebhookVerifierToken(v string)`

SetWebhookVerifierToken sets WebhookVerifierToken field to given value.

### HasWebhookVerifierToken

`func (o *TypesQuickBooksConnectionMetadata) HasWebhookVerifierToken() bool`

HasWebhookVerifierToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


