# Release History

## 2.0.0 (2023-04-28)
### Breaking Changes

- Type of `ServiceProperties.ProvisioningState` has been changed from `*ProvisioningState` to `*CommunicationServicesProvisioningState`
- Function `NewServiceClient` has been removed
- Function `*ServiceClient.CheckNameAvailability` has been removed
- Function `*ServiceClient.BeginCreateOrUpdate` has been removed
- Function `*ServiceClient.BeginDelete` has been removed
- Function `*ServiceClient.Get` has been removed
- Function `*ServiceClient.LinkNotificationHub` has been removed
- Function `*ServiceClient.NewListByResourceGroupPager` has been removed
- Function `*ServiceClient.NewListBySubscriptionPager` has been removed
- Function `*ServiceClient.ListKeys` has been removed
- Function `*ServiceClient.RegenerateKey` has been removed
- Function `*ServiceClient.Update` has been removed
- Struct `LocationResource` has been removed
- Struct `NameAvailability` has been removed
- Struct `ServiceClient` has been removed

### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module
- New enum type `CheckNameAvailabilityReason` with values `CheckNameAvailabilityReasonAlreadyExists`, `CheckNameAvailabilityReasonInvalid`
- New enum type `CommunicationServicesProvisioningState` with values `CommunicationServicesProvisioningStateCanceled`, `CommunicationServicesProvisioningStateCreating`, `CommunicationServicesProvisioningStateDeleting`, `CommunicationServicesProvisioningStateFailed`, `CommunicationServicesProvisioningStateMoving`, `CommunicationServicesProvisioningStateRunning`, `CommunicationServicesProvisioningStateSucceeded`, `CommunicationServicesProvisioningStateUnknown`, `CommunicationServicesProvisioningStateUpdating`
- New enum type `DomainManagement` with values `DomainManagementAzureManaged`, `DomainManagementCustomerManaged`, `DomainManagementCustomerManagedInExchangeOnline`
- New enum type `DomainsProvisioningState` with values `DomainsProvisioningStateCanceled`, `DomainsProvisioningStateCreating`, `DomainsProvisioningStateDeleting`, `DomainsProvisioningStateFailed`, `DomainsProvisioningStateMoving`, `DomainsProvisioningStateRunning`, `DomainsProvisioningStateSucceeded`, `DomainsProvisioningStateUnknown`, `DomainsProvisioningStateUpdating`
- New enum type `EmailServicesProvisioningState` with values `EmailServicesProvisioningStateCanceled`, `EmailServicesProvisioningStateCreating`, `EmailServicesProvisioningStateDeleting`, `EmailServicesProvisioningStateFailed`, `EmailServicesProvisioningStateMoving`, `EmailServicesProvisioningStateRunning`, `EmailServicesProvisioningStateSucceeded`, `EmailServicesProvisioningStateUnknown`, `EmailServicesProvisioningStateUpdating`
- New enum type `UserEngagementTracking` with values `UserEngagementTrackingDisabled`, `UserEngagementTrackingEnabled`
- New enum type `VerificationStatus` with values `VerificationStatusCancellationRequested`, `VerificationStatusNotStarted`, `VerificationStatusVerificationFailed`, `VerificationStatusVerificationInProgress`, `VerificationStatusVerificationRequested`, `VerificationStatusVerified`
- New enum type `VerificationType` with values `VerificationTypeDKIM`, `VerificationTypeDKIM2`, `VerificationTypeDMARC`, `VerificationTypeDomain`, `VerificationTypeSPF`
- New function `NewDomainsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*DomainsClient, error)`
- New function `*DomainsClient.BeginCancelVerification(context.Context, string, string, string, VerificationParameter, *DomainsClientBeginCancelVerificationOptions) (*runtime.Poller[DomainsClientCancelVerificationResponse], error)`
- New function `*DomainsClient.BeginCreateOrUpdate(context.Context, string, string, string, DomainResource, *DomainsClientBeginCreateOrUpdateOptions) (*runtime.Poller[DomainsClientCreateOrUpdateResponse], error)`
- New function `*DomainsClient.BeginDelete(context.Context, string, string, string, *DomainsClientBeginDeleteOptions) (*runtime.Poller[DomainsClientDeleteResponse], error)`
- New function `*DomainsClient.Get(context.Context, string, string, string, *DomainsClientGetOptions) (DomainsClientGetResponse, error)`
- New function `*DomainsClient.BeginInitiateVerification(context.Context, string, string, string, VerificationParameter, *DomainsClientBeginInitiateVerificationOptions) (*runtime.Poller[DomainsClientInitiateVerificationResponse], error)`
- New function `*DomainsClient.NewListByEmailServiceResourcePager(string, string, *DomainsClientListByEmailServiceResourceOptions) *runtime.Pager[DomainsClientListByEmailServiceResourceResponse]`
- New function `*DomainsClient.BeginUpdate(context.Context, string, string, string, UpdateDomainRequestParameters, *DomainsClientBeginUpdateOptions) (*runtime.Poller[DomainsClientUpdateResponse], error)`
- New function `NewEmailServicesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*EmailServicesClient, error)`
- New function `*EmailServicesClient.BeginCreateOrUpdate(context.Context, string, string, EmailServiceResource, *EmailServicesClientBeginCreateOrUpdateOptions) (*runtime.Poller[EmailServicesClientCreateOrUpdateResponse], error)`
- New function `*EmailServicesClient.BeginDelete(context.Context, string, string, *EmailServicesClientBeginDeleteOptions) (*runtime.Poller[EmailServicesClientDeleteResponse], error)`
- New function `*EmailServicesClient.Get(context.Context, string, string, *EmailServicesClientGetOptions) (EmailServicesClientGetResponse, error)`
- New function `*EmailServicesClient.NewListByResourceGroupPager(string, *EmailServicesClientListByResourceGroupOptions) *runtime.Pager[EmailServicesClientListByResourceGroupResponse]`
- New function `*EmailServicesClient.NewListBySubscriptionPager(*EmailServicesClientListBySubscriptionOptions) *runtime.Pager[EmailServicesClientListBySubscriptionResponse]`
- New function `*EmailServicesClient.ListVerifiedExchangeOnlineDomains(context.Context, *EmailServicesClientListVerifiedExchangeOnlineDomainsOptions) (EmailServicesClientListVerifiedExchangeOnlineDomainsResponse, error)`
- New function `*EmailServicesClient.BeginUpdate(context.Context, string, string, EmailServiceResourceUpdate, *EmailServicesClientBeginUpdateOptions) (*runtime.Poller[EmailServicesClientUpdateResponse], error)`
- New function `NewSenderUsernamesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SenderUsernamesClient, error)`
- New function `*SenderUsernamesClient.CreateOrUpdate(context.Context, string, string, string, string, SenderUsernameResource, *SenderUsernamesClientCreateOrUpdateOptions) (SenderUsernamesClientCreateOrUpdateResponse, error)`
- New function `*SenderUsernamesClient.Delete(context.Context, string, string, string, string, *SenderUsernamesClientDeleteOptions) (SenderUsernamesClientDeleteResponse, error)`
- New function `*SenderUsernamesClient.Get(context.Context, string, string, string, string, *SenderUsernamesClientGetOptions) (SenderUsernamesClientGetResponse, error)`
- New function `*SenderUsernamesClient.NewListByDomainsPager(string, string, string, *SenderUsernamesClientListByDomainsOptions) *runtime.Pager[SenderUsernamesClientListByDomainsResponse]`
- New function `NewServicesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ServicesClient, error)`
- New function `*ServicesClient.CheckNameAvailability(context.Context, NameAvailabilityParameters, *ServicesClientCheckNameAvailabilityOptions) (ServicesClientCheckNameAvailabilityResponse, error)`
- New function `*ServicesClient.BeginCreateOrUpdate(context.Context, string, string, ServiceResource, *ServicesClientBeginCreateOrUpdateOptions) (*runtime.Poller[ServicesClientCreateOrUpdateResponse], error)`
- New function `*ServicesClient.BeginDelete(context.Context, string, string, *ServicesClientBeginDeleteOptions) (*runtime.Poller[ServicesClientDeleteResponse], error)`
- New function `*ServicesClient.Get(context.Context, string, string, *ServicesClientGetOptions) (ServicesClientGetResponse, error)`
- New function `*ServicesClient.LinkNotificationHub(context.Context, string, string, *ServicesClientLinkNotificationHubOptions) (ServicesClientLinkNotificationHubResponse, error)`
- New function `*ServicesClient.NewListByResourceGroupPager(string, *ServicesClientListByResourceGroupOptions) *runtime.Pager[ServicesClientListByResourceGroupResponse]`
- New function `*ServicesClient.NewListBySubscriptionPager(*ServicesClientListBySubscriptionOptions) *runtime.Pager[ServicesClientListBySubscriptionResponse]`
- New function `*ServicesClient.ListKeys(context.Context, string, string, *ServicesClientListKeysOptions) (ServicesClientListKeysResponse, error)`
- New function `*ServicesClient.RegenerateKey(context.Context, string, string, RegenerateKeyParameters, *ServicesClientRegenerateKeyOptions) (ServicesClientRegenerateKeyResponse, error)`
- New function `*ServicesClient.Update(context.Context, string, string, ServiceResourceUpdate, *ServicesClientUpdateOptions) (ServicesClientUpdateResponse, error)`
- New struct `CheckNameAvailabilityRequest`
- New struct `CheckNameAvailabilityResponse`
- New struct `DNSRecord`
- New struct `DomainProperties`
- New struct `DomainPropertiesVerificationRecords`
- New struct `DomainPropertiesVerificationStates`
- New struct `DomainResource`
- New struct `DomainResourceList`
- New struct `EmailServiceProperties`
- New struct `EmailServiceResource`
- New struct `EmailServiceResourceList`
- New struct `EmailServiceResourceUpdate`
- New struct `ProxyResource`
- New struct `SenderUsernameProperties`
- New struct `SenderUsernameResource`
- New struct `SenderUsernameResourceCollection`
- New struct `ServiceResourceUpdate`
- New struct `ServiceUpdateProperties`
- New struct `TrackedResource`
- New struct `UpdateDomainProperties`
- New struct `UpdateDomainRequestParameters`
- New struct `VerificationParameter`
- New struct `VerificationStatusRecord`
- New field `SystemData` in struct `Resource`
- New field `LinkedDomains` in struct `ServiceProperties`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/communication/armcommunication` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).
