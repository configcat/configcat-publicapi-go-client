# CodeReferenceRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigId** | **string** | The Config&#x27;s identifier the scanning was performed against. | [default to null]
**Repository** | **string** | The source control repository that contains the scanned code. (Source of the repository selector on the ConfigCat Dashboard) | [default to null]
**Branch** | **string** | The source control branch on where the scan was performed. (Source of the branch selector on the ConfigCat Dashboard) | [default to null]
**CommitUrl** | **string** | The related commit&#x27;s URL. (Appears on the ConfigCat Dashboard) | [optional] [default to null]
**CommitHash** | **string** | The related commit&#x27;s hash. (Appears on the ConfigCat Dashboard) | [optional] [default to null]
**Uploader** | **string** | The scanning tool&#x27;s name. (Appears on the ConfigCat Dashboard) | [optional] [default to null]
**ActiveBranches** | **[]string** | The currently active branches of the repository. Each previously uploaded report that belongs to a non-reported active branch is being deleted. | [optional] [default to null]
**FlagReferences** | [**[]FlagReference**](FlagReference.md) | The actual code reference collection. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

