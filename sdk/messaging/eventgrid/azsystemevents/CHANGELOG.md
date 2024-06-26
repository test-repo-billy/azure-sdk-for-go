# Release History

## 0.3.1 (Unreleased)

### Features Added

### Breaking Changes

### Bugs Fixed

### Other Changes

## 0.3.0 (2024-04-03)

### Features Added

- Added events ACSRouterWorkerUpdatedEventData and ACSAdvancedMessageDeliveryStatusUpdatedEventData. (PR#22638)

### Breaking Changes

Field and type renames:
- Globally, types and fields named ChannelType has been renamed to ChannelKind
- ACS events and constants have been changed to use an all-caps name (ex: AcsEmailDeliveryReportStatusDetails -> ACSEmailDeliveryReportStatusDetails).
- ACSAdvancedMessageContext.ID -> MessageID
- ACSAdvancedMessageReceivedEventData
  - .Media -> MediaContent
  - .Interactive -> InteractiveContent

## 0.2.0 (2024-03-14)

### Features Added

- Added API Center system events under their official names.

### Breaking Changes

- Events have been renamed:
  - APIDefinitionAddedEventData renamed to APICenterAPIDefinitionAddedEventData
  - APIDefinitionUpdatedEventData renamed to APICenterAPIDefinitionUpdatedEventData

## 0.1.0 (2024-03-05)

### Features Added

- Initial preview for Event Grid system events.
