# InlineObject4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Library** | **string** | The library in which the audio should be stored | [optional] 
**ImportReference** | **string** |  | [optional] 
**Source** | **string** |  | [optional] 
**AudioFile** | [***os.File**](*os.File.md) |  | [optional] 
**ImportStatus** | **string** | Setting import_status to draft will prevent processing, but allow further modifications to audio and metadata. Once ready, use the PATCH method to set import_status to pending. Default to &#x60;pending&#x60; if unspecified. | [optional] [default to IMPORT_STATUS_PENDING]
**ImportMetadata** | [**ImportMetadata**](ImportMetadata.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


