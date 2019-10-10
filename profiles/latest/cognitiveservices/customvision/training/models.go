// +build go1.9

// Copyright 2019 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package training

import original "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v3.1/customvision/training"

type Classifier = original.Classifier

const (
	Multiclass Classifier = original.Multiclass
	Multilabel Classifier = original.Multilabel
)

type CustomVisionErrorCodes = original.CustomVisionErrorCodes

const (
	BadRequest                                                  CustomVisionErrorCodes = original.BadRequest
	BadRequestCannotMigrateProjectWithName                      CustomVisionErrorCodes = original.BadRequestCannotMigrateProjectWithName
	BadRequestClassificationTrainingValidationFailed            CustomVisionErrorCodes = original.BadRequestClassificationTrainingValidationFailed
	BadRequestDetectionTrainingNotAllowNegativeTag              CustomVisionErrorCodes = original.BadRequestDetectionTrainingNotAllowNegativeTag
	BadRequestDetectionTrainingValidationFailed                 CustomVisionErrorCodes = original.BadRequestDetectionTrainingValidationFailed
	BadRequestDomainNotSupportedForAdvancedTraining             CustomVisionErrorCodes = original.BadRequestDomainNotSupportedForAdvancedTraining
	BadRequestExceededBatchSize                                 CustomVisionErrorCodes = original.BadRequestExceededBatchSize
	BadRequestExceededQuota                                     CustomVisionErrorCodes = original.BadRequestExceededQuota
	BadRequestExceedIterationPerProjectLimit                    CustomVisionErrorCodes = original.BadRequestExceedIterationPerProjectLimit
	BadRequestExceedProjectLimit                                CustomVisionErrorCodes = original.BadRequestExceedProjectLimit
	BadRequestExceedTagPerImageLimit                            CustomVisionErrorCodes = original.BadRequestExceedTagPerImageLimit
	BadRequestExceedTagPerProjectLimit                          CustomVisionErrorCodes = original.BadRequestExceedTagPerProjectLimit
	BadRequestExportAlreadyInProgress                           CustomVisionErrorCodes = original.BadRequestExportAlreadyInProgress
	BadRequestExportPlatformNotSupportedForAdvancedTraining     CustomVisionErrorCodes = original.BadRequestExportPlatformNotSupportedForAdvancedTraining
	BadRequestExportValidationFailed                            CustomVisionErrorCodes = original.BadRequestExportValidationFailed
	BadRequestImageBatch                                        CustomVisionErrorCodes = original.BadRequestImageBatch
	BadRequestImageExceededCount                                CustomVisionErrorCodes = original.BadRequestImageExceededCount
	BadRequestImageFormat                                       CustomVisionErrorCodes = original.BadRequestImageFormat
	BadRequestImageRegions                                      CustomVisionErrorCodes = original.BadRequestImageRegions
	BadRequestImageSizeBytes                                    CustomVisionErrorCodes = original.BadRequestImageSizeBytes
	BadRequestImageStream                                       CustomVisionErrorCodes = original.BadRequestImageStream
	BadRequestImageTags                                         CustomVisionErrorCodes = original.BadRequestImageTags
	BadRequestImageURL                                          CustomVisionErrorCodes = original.BadRequestImageURL
	BadRequestInvalid                                           CustomVisionErrorCodes = original.BadRequestInvalid
	BadRequestInvalidEmailAddress                               CustomVisionErrorCodes = original.BadRequestInvalidEmailAddress
	BadRequestInvalidIds                                        CustomVisionErrorCodes = original.BadRequestInvalidIds
	BadRequestInvalidPublishName                                CustomVisionErrorCodes = original.BadRequestInvalidPublishName
	BadRequestInvalidPublishTarget                              CustomVisionErrorCodes = original.BadRequestInvalidPublishTarget
	BadRequestIterationDescription                              CustomVisionErrorCodes = original.BadRequestIterationDescription
	BadRequestIterationIsNotTrained                             CustomVisionErrorCodes = original.BadRequestIterationIsNotTrained
	BadRequestIterationIsPublished                              CustomVisionErrorCodes = original.BadRequestIterationIsPublished
	BadRequestIterationName                                     CustomVisionErrorCodes = original.BadRequestIterationName
	BadRequestIterationNameNotUnique                            CustomVisionErrorCodes = original.BadRequestIterationNameNotUnique
	BadRequestIterationNotPublished                             CustomVisionErrorCodes = original.BadRequestIterationNotPublished
	BadRequestMultiClassClassificationTrainingValidationFailed  CustomVisionErrorCodes = original.BadRequestMultiClassClassificationTrainingValidationFailed
	BadRequestMultiLabelClassificationTrainingValidationFailed  CustomVisionErrorCodes = original.BadRequestMultiLabelClassificationTrainingValidationFailed
	BadRequestMultipleNegativeTag                               CustomVisionErrorCodes = original.BadRequestMultipleNegativeTag
	BadRequestNegativeAndRegularTagOnSameImage                  CustomVisionErrorCodes = original.BadRequestNegativeAndRegularTagOnSameImage
	BadRequestNotLimitedTrial                                   CustomVisionErrorCodes = original.BadRequestNotLimitedTrial
	BadRequestNotSupported                                      CustomVisionErrorCodes = original.BadRequestNotSupported
	BadRequestPredictionIdsExceededCount                        CustomVisionErrorCodes = original.BadRequestPredictionIdsExceededCount
	BadRequestPredictionIdsMissing                              CustomVisionErrorCodes = original.BadRequestPredictionIdsMissing
	BadRequestPredictionInvalidApplicationName                  CustomVisionErrorCodes = original.BadRequestPredictionInvalidApplicationName
	BadRequestPredictionInvalidQueryParameters                  CustomVisionErrorCodes = original.BadRequestPredictionInvalidQueryParameters
	BadRequestPredictionResultsExceededCount                    CustomVisionErrorCodes = original.BadRequestPredictionResultsExceededCount
	BadRequestPredictionTagsExceededCount                       CustomVisionErrorCodes = original.BadRequestPredictionTagsExceededCount
	BadRequestProjectDescription                                CustomVisionErrorCodes = original.BadRequestProjectDescription
	BadRequestProjectImagePreprocessingSettings                 CustomVisionErrorCodes = original.BadRequestProjectImagePreprocessingSettings
	BadRequestProjectName                                       CustomVisionErrorCodes = original.BadRequestProjectName
	BadRequestProjectNameNotUnique                              CustomVisionErrorCodes = original.BadRequestProjectNameNotUnique
	BadRequestProjectUnknownClassification                      CustomVisionErrorCodes = original.BadRequestProjectUnknownClassification
	BadRequestProjectUnknownDomain                              CustomVisionErrorCodes = original.BadRequestProjectUnknownDomain
	BadRequestProjectUnsupportedDomainTypeChange                CustomVisionErrorCodes = original.BadRequestProjectUnsupportedDomainTypeChange
	BadRequestProjectUnsupportedExportPlatform                  CustomVisionErrorCodes = original.BadRequestProjectUnsupportedExportPlatform
	BadRequestRequiredParamIsNull                               CustomVisionErrorCodes = original.BadRequestRequiredParamIsNull
	BadRequestReservedBudgetInHoursNotEnoughForAdvancedTraining CustomVisionErrorCodes = original.BadRequestReservedBudgetInHoursNotEnoughForAdvancedTraining
	BadRequestSubscriptionAPI                                   CustomVisionErrorCodes = original.BadRequestSubscriptionAPI
	BadRequestTagDescription                                    CustomVisionErrorCodes = original.BadRequestTagDescription
	BadRequestTagName                                           CustomVisionErrorCodes = original.BadRequestTagName
	BadRequestTagNameNotUnique                                  CustomVisionErrorCodes = original.BadRequestTagNameNotUnique
	BadRequestTagType                                           CustomVisionErrorCodes = original.BadRequestTagType
	BadRequestTrainingAlreadyInProgress                         CustomVisionErrorCodes = original.BadRequestTrainingAlreadyInProgress
	BadRequestTrainingNotNeeded                                 CustomVisionErrorCodes = original.BadRequestTrainingNotNeeded
	BadRequestTrainingNotNeededButTrainingPipelineUpdated       CustomVisionErrorCodes = original.BadRequestTrainingNotNeededButTrainingPipelineUpdated
	BadRequestTrainingValidationFailed                          CustomVisionErrorCodes = original.BadRequestTrainingValidationFailed
	BadRequestUnpublishFailed                                   CustomVisionErrorCodes = original.BadRequestUnpublishFailed
	BadRequestWorkspaceCannotBeModified                         CustomVisionErrorCodes = original.BadRequestWorkspaceCannotBeModified
	BadRequestWorkspaceNotDeletable                             CustomVisionErrorCodes = original.BadRequestWorkspaceNotDeletable
	Conflict                                                    CustomVisionErrorCodes = original.Conflict
	ConflictInvalid                                             CustomVisionErrorCodes = original.ConflictInvalid
	ErrorExporterInvalidClassifier                              CustomVisionErrorCodes = original.ErrorExporterInvalidClassifier
	ErrorExporterInvalidFeaturizer                              CustomVisionErrorCodes = original.ErrorExporterInvalidFeaturizer
	ErrorExporterInvalidPlatform                                CustomVisionErrorCodes = original.ErrorExporterInvalidPlatform
	ErrorFeaturizationAugmentationError                         CustomVisionErrorCodes = original.ErrorFeaturizationAugmentationError
	ErrorFeaturizationAugmentationUnavailable                   CustomVisionErrorCodes = original.ErrorFeaturizationAugmentationUnavailable
	ErrorFeaturizationInvalidFeaturizer                         CustomVisionErrorCodes = original.ErrorFeaturizationInvalidFeaturizer
	ErrorFeaturizationQueueTimeout                              CustomVisionErrorCodes = original.ErrorFeaturizationQueueTimeout
	ErrorFeaturizationServiceUnavailable                        CustomVisionErrorCodes = original.ErrorFeaturizationServiceUnavailable
	ErrorFeaturizationUnrecognizedJob                           CustomVisionErrorCodes = original.ErrorFeaturizationUnrecognizedJob
	ErrorInvalid                                                CustomVisionErrorCodes = original.ErrorInvalid
	ErrorPrediction                                             CustomVisionErrorCodes = original.ErrorPrediction
	ErrorPredictionModelNotCached                               CustomVisionErrorCodes = original.ErrorPredictionModelNotCached
	ErrorPredictionModelNotFound                                CustomVisionErrorCodes = original.ErrorPredictionModelNotFound
	ErrorPredictionServiceUnavailable                           CustomVisionErrorCodes = original.ErrorPredictionServiceUnavailable
	ErrorPredictionStorage                                      CustomVisionErrorCodes = original.ErrorPredictionStorage
	ErrorProjectExportRequestFailed                             CustomVisionErrorCodes = original.ErrorProjectExportRequestFailed
	ErrorProjectInvalidDomain                                   CustomVisionErrorCodes = original.ErrorProjectInvalidDomain
	ErrorProjectInvalidPipelineConfiguration                    CustomVisionErrorCodes = original.ErrorProjectInvalidPipelineConfiguration
	ErrorProjectInvalidWorkspace                                CustomVisionErrorCodes = original.ErrorProjectInvalidWorkspace
	ErrorProjectTrainingRequestFailed                           CustomVisionErrorCodes = original.ErrorProjectTrainingRequestFailed
	ErrorRegionProposal                                         CustomVisionErrorCodes = original.ErrorRegionProposal
	ErrorUnknown                                                CustomVisionErrorCodes = original.ErrorUnknown
	Forbidden                                                   CustomVisionErrorCodes = original.Forbidden
	ForbiddenDRModeEnabled                                      CustomVisionErrorCodes = original.ForbiddenDRModeEnabled
	ForbiddenInvalid                                            CustomVisionErrorCodes = original.ForbiddenInvalid
	ForbiddenUser                                               CustomVisionErrorCodes = original.ForbiddenUser
	ForbiddenUserDisabled                                       CustomVisionErrorCodes = original.ForbiddenUserDisabled
	ForbiddenUserDoesNotExist                                   CustomVisionErrorCodes = original.ForbiddenUserDoesNotExist
	ForbiddenUserInsufficientCapability                         CustomVisionErrorCodes = original.ForbiddenUserInsufficientCapability
	ForbiddenUserResource                                       CustomVisionErrorCodes = original.ForbiddenUserResource
	ForbiddenUserSignupAllowanceExceeded                        CustomVisionErrorCodes = original.ForbiddenUserSignupAllowanceExceeded
	ForbiddenUserSignupDisabled                                 CustomVisionErrorCodes = original.ForbiddenUserSignupDisabled
	NoError                                                     CustomVisionErrorCodes = original.NoError
	NotFound                                                    CustomVisionErrorCodes = original.NotFound
	NotFoundApimSubscription                                    CustomVisionErrorCodes = original.NotFoundApimSubscription
	NotFoundDomain                                              CustomVisionErrorCodes = original.NotFoundDomain
	NotFoundImage                                               CustomVisionErrorCodes = original.NotFoundImage
	NotFoundInvalid                                             CustomVisionErrorCodes = original.NotFoundInvalid
	NotFoundIteration                                           CustomVisionErrorCodes = original.NotFoundIteration
	NotFoundIterationPerformance                                CustomVisionErrorCodes = original.NotFoundIterationPerformance
	NotFoundProject                                             CustomVisionErrorCodes = original.NotFoundProject
	NotFoundProjectDefaultIteration                             CustomVisionErrorCodes = original.NotFoundProjectDefaultIteration
	NotFoundTag                                                 CustomVisionErrorCodes = original.NotFoundTag
	UnsupportedMediaType                                        CustomVisionErrorCodes = original.UnsupportedMediaType
)

type DomainType = original.DomainType

const (
	Classification  DomainType = original.Classification
	ObjectDetection DomainType = original.ObjectDetection
)

type ExportFlavor = original.ExportFlavor

const (
	ARM              ExportFlavor = original.ARM
	Linux            ExportFlavor = original.Linux
	ONNX10           ExportFlavor = original.ONNX10
	ONNX12           ExportFlavor = original.ONNX12
	TensorFlowLite   ExportFlavor = original.TensorFlowLite
	TensorFlowNormal ExportFlavor = original.TensorFlowNormal
	Windows          ExportFlavor = original.Windows
)

type ExportPlatform = original.ExportPlatform

const (
	CoreML     ExportPlatform = original.CoreML
	DockerFile ExportPlatform = original.DockerFile
	ONNX       ExportPlatform = original.ONNX
	TensorFlow ExportPlatform = original.TensorFlow
	VAIDK      ExportPlatform = original.VAIDK
)

type ExportStatus = original.ExportStatus

const (
	Done      ExportStatus = original.Done
	Exporting ExportStatus = original.Exporting
	Failed    ExportStatus = original.Failed
)

type ImageCreateStatus = original.ImageCreateStatus

const (
	ImageCreateStatusErrorImageFormat                      ImageCreateStatus = original.ImageCreateStatusErrorImageFormat
	ImageCreateStatusErrorImageSize                        ImageCreateStatus = original.ImageCreateStatusErrorImageSize
	ImageCreateStatusErrorLimitExceed                      ImageCreateStatus = original.ImageCreateStatusErrorLimitExceed
	ImageCreateStatusErrorNegativeAndRegularTagOnSameImage ImageCreateStatus = original.ImageCreateStatusErrorNegativeAndRegularTagOnSameImage
	ImageCreateStatusErrorRegionLimitExceed                ImageCreateStatus = original.ImageCreateStatusErrorRegionLimitExceed
	ImageCreateStatusErrorSource                           ImageCreateStatus = original.ImageCreateStatusErrorSource
	ImageCreateStatusErrorStorage                          ImageCreateStatus = original.ImageCreateStatusErrorStorage
	ImageCreateStatusErrorTagLimitExceed                   ImageCreateStatus = original.ImageCreateStatusErrorTagLimitExceed
	ImageCreateStatusErrorUnknown                          ImageCreateStatus = original.ImageCreateStatusErrorUnknown
	ImageCreateStatusOK                                    ImageCreateStatus = original.ImageCreateStatusOK
	ImageCreateStatusOKDuplicate                           ImageCreateStatus = original.ImageCreateStatusOKDuplicate
)

type OrderBy = original.OrderBy

const (
	Newest    OrderBy = original.Newest
	Oldest    OrderBy = original.Oldest
	Suggested OrderBy = original.Suggested
)

type SortBy = original.SortBy

const (
	UncertaintyAscending  SortBy = original.UncertaintyAscending
	UncertaintyDescending SortBy = original.UncertaintyDescending
)

type TagType = original.TagType

const (
	Negative TagType = original.Negative
	Regular  TagType = original.Regular
)

type Type = original.Type

const (
	TypeAdvanced Type = original.TypeAdvanced
	TypeRegular  Type = original.TypeRegular
)

type BaseClient = original.BaseClient
type Bool = original.Bool
type BoundingBox = original.BoundingBox
type CustomVisionError = original.CustomVisionError
type Domain = original.Domain
type Export = original.Export
type Image = original.Image
type ImageCreateResult = original.ImageCreateResult
type ImageCreateSummary = original.ImageCreateSummary
type ImageFileCreateBatch = original.ImageFileCreateBatch
type ImageFileCreateEntry = original.ImageFileCreateEntry
type ImageIDCreateBatch = original.ImageIDCreateBatch
type ImageIDCreateEntry = original.ImageIDCreateEntry
type ImagePerformance = original.ImagePerformance
type ImagePrediction = original.ImagePrediction
type ImageProcessingSettings = original.ImageProcessingSettings
type ImageRegion = original.ImageRegion
type ImageRegionCreateBatch = original.ImageRegionCreateBatch
type ImageRegionCreateEntry = original.ImageRegionCreateEntry
type ImageRegionCreateResult = original.ImageRegionCreateResult
type ImageRegionCreateSummary = original.ImageRegionCreateSummary
type ImageRegionProposal = original.ImageRegionProposal
type ImageTag = original.ImageTag
type ImageTagCreateBatch = original.ImageTagCreateBatch
type ImageTagCreateEntry = original.ImageTagCreateEntry
type ImageTagCreateSummary = original.ImageTagCreateSummary
type ImageURL = original.ImageURL
type ImageURLCreateBatch = original.ImageURLCreateBatch
type ImageURLCreateEntry = original.ImageURLCreateEntry
type Int32 = original.Int32
type Iteration = original.Iteration
type IterationPerformance = original.IterationPerformance
type ListDomain = original.ListDomain
type ListExport = original.ListExport
type ListImage = original.ListImage
type ListImagePerformance = original.ListImagePerformance
type ListIteration = original.ListIteration
type ListProject = original.ListProject
type ListSuggestedTagAndRegion = original.ListSuggestedTagAndRegion
type ListTag = original.ListTag
type Prediction = original.Prediction
type PredictionQueryResult = original.PredictionQueryResult
type PredictionQueryTag = original.PredictionQueryTag
type PredictionQueryToken = original.PredictionQueryToken
type Project = original.Project
type ProjectSettings = original.ProjectSettings
type Region = original.Region
type RegionProposal = original.RegionProposal
type SetInt32 = original.SetInt32
type StoredImagePrediction = original.StoredImagePrediction
type StoredSuggestedTagAndRegion = original.StoredSuggestedTagAndRegion
type SuggestedTagAndRegion = original.SuggestedTagAndRegion
type SuggestedTagAndRegionQuery = original.SuggestedTagAndRegionQuery
type SuggestedTagAndRegionQueryToken = original.SuggestedTagAndRegionQueryToken
type Tag = original.Tag
type TagFilter = original.TagFilter
type TagPerformance = original.TagPerformance

func New(aPIKey string, endpoint string) BaseClient {
	return original.New(aPIKey, endpoint)
}
func NewWithoutDefaults(aPIKey string, endpoint string) BaseClient {
	return original.NewWithoutDefaults(aPIKey, endpoint)
}
func PossibleClassifierValues() []Classifier {
	return original.PossibleClassifierValues()
}
func PossibleCustomVisionErrorCodesValues() []CustomVisionErrorCodes {
	return original.PossibleCustomVisionErrorCodesValues()
}
func PossibleDomainTypeValues() []DomainType {
	return original.PossibleDomainTypeValues()
}
func PossibleExportFlavorValues() []ExportFlavor {
	return original.PossibleExportFlavorValues()
}
func PossibleExportPlatformValues() []ExportPlatform {
	return original.PossibleExportPlatformValues()
}
func PossibleExportStatusValues() []ExportStatus {
	return original.PossibleExportStatusValues()
}
func PossibleImageCreateStatusValues() []ImageCreateStatus {
	return original.PossibleImageCreateStatusValues()
}
func PossibleOrderByValues() []OrderBy {
	return original.PossibleOrderByValues()
}
func PossibleSortByValues() []SortBy {
	return original.PossibleSortByValues()
}
func PossibleTagTypeValues() []TagType {
	return original.PossibleTagTypeValues()
}
func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
