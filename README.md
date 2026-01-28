# GOMP

Open Multiplayer C-API implementation for Golang.

Under development. Do not use in production under any circumstances.

C-API: https://github.com/openmultiplayer/open.mp-capi

### API

API wrappers.


- [x] **Player**

    - [x] SetSpawnInfo

    - [x] GetSpawnInfo

    - [x] GetNetworkStats

    - [x] NetStatsBytesReceived

    - [x] NetStatsBytesSent

    - [x] NetStatsConnectionStatus

    - [x] NetStatsGetConnectedTime

    - [x] NetStatsGetIpPort

    - [x] NetStatsMessagesReceived

    - [x] NetStatsMessagesRecvPerSecond

    - [x] NetStatsMessagesSent

    - [x] NetStatsPacketLossPercent

    - [x] GetCustomSkin

    - [x] GetDialog

    - [x] GetDialogData

    - [x] GetMenu

    - [x] GetSurfingPlayerObject

    - [x] GetCameraTargetPlayerObject

    - [x] FromID

    - [x] GetID

    - [x] SendClientMessage

    - [x] SetCameraPos

    - [x] SetDrunkLevel

    - [x] SetInterior

    - [x] SetWantedLevel

    - [x] SetWeather

    - [x] GetWeather

    - [x] SetSkin

    - [x] SetShopName

    - [x] GiveMoney

    - [x] SetCameraLookAt

    - [x] SetCameraBehind

    - [x] CreateExplosion

    - [x] PlayAudioStream

    - [x] StopAudioStream

    - [x] ToggleWidescreen

    - [x] IsWidescreenToggled

    - [x] SetHealth

    - [x] GetHealth

    - [x] SetArmor

    - [x] GetArmor

    - [x] SetTeam

    - [x] GetTeam

    - [x] SetScore

    - [x] GetScore

    - [x] GetSkin

    - [x] SetColor

    - [x] GetColor

    - [x] GetDefaultColor

    - [x] GetDrunkLevel

    - [x] GiveWeapon

    - [x] RemoveWeapon

    - [x] GetMoney

    - [x] ResetMoney

    - [x] SetName

    - [x] GetName

    - [x] GetState

    - [x] GetPing

    - [x] GetWeapon

    - [x] SetTime

    - [x] GetTime

    - [x] ToggleClock

    - [x] HasClock

    - [x] ForceClassSelection

    - [x] GetWantedLevel

    - [x] SetFightingStyle

    - [x] GetFightingStyle

    - [x] SetVelocity

    - [x] GetVelocity

    - [x] GetCameraPos

    - [x] GetDistanceFromPoint

    - [x] GetInterior

    - [x] SetPos

    - [x] GetPos

    - [x] GetVirtualWorld

    - [x] IsNPC

    - [x] IsStreamedIn

    - [x] PlayGameSound

    - [x] SpectatePlayer

    - [x] SpectateVehicle

    - [x] SetVirtualWorld

    - [x] SetWorldBounds

    - [x] ClearWorldBounds

    - [x] GetWorldBounds

    - [x] ClearAnimations

    - [x] GetLastShotVectors

    - [x] GetCameraTargetPlayer

    - [x] GetCameraTargetActor

    - [x] GetCameraTargetObject

    - [x] GetCameraTargetVehicle

    - [x] PutInVehicle

    - [x] RemoveBuilding

    - [x] GetBuildingsRemoved

    - [x] RemoveFromVehicle

    - [x] RemoveMapIcon

    - [x] SetMapIcon

    - [x] ResetWeapons

    - [x] SetAmmo

    - [x] SetArmedWeapon

    - [x] SetChatBubble

    - [x] SetPosFindZ

    - [x] SetSkillLevel

    - [x] SetSpecialAction

    - [x] ShowNameTagForPlayer

    - [x] ToggleControllable

    - [x] ToggleSpectating

    - [x] ApplyAnimation

    - [x] GetAnimationName

    - [x] EditAttachedObject

    - [x] EnableCameraTarget

    - [x] EnableStuntBonus

    - [x] GetPlayerAmmo

    - [x] GetAnimationIndex

    - [x] GetFacingAngle

    - [x] GetIp

    - [x] GetSpecialAction

    - [x] GetVehicleID

    - [x] GetVehicleSeat

    - [x] GetWeaponData

    - [x] GetWeaponState

    - [x] InterpolateCameraPos

    - [x] InterpolateCameraLookAt

    - [x] IsPlayerAttachedObjectSlotUsed

    - [x] AttachCameraToObject

    - [x] AttachCameraToPlayerObject

    - [x] GetCameraAspectRatio

    - [x] GetCameraFrontVector

    - [x] GetCameraMode

    - [x] GetKeys

    - [x] GetSurfingVehicle

    - [x] GetSurfingObject

    - [x] GetTargetPlayer

    - [x] GetTargetActor

    - [x] IsInVehicle

    - [x] IsInAnyVehicle

    - [x] IsInRangeOfPoint

    - [x] PlayCrimeReport

    - [x] RemoveAttachedObject

    - [x] SetAttachedObject

    - [x] GetAttachedObject

    - [x] SetFacingAngle

    - [x] SetMarkerForPlayer

    - [x] GetMarkerForPlayer

    - [x] AllowTeleport

    - [x] IsTeleportAllowed

    - [x] DisableRemoteVehicleCollisions

    - [x] GetCameraZoom

    - [x] SelectTextDraw

    - [x] CancelSelectTextDraw

    - [x] SendClientCheck

    - [x] Spawn

    - [x] GPCI

    - [x] IsAdmin

    - [x] Kick

    - [x] ShowGameText

    - [x] HideGameText

    - [x] HasGameText

    - [x] GetGameText

    - [x] Ban

    - [x] BanEx

    - [x] SendDeathMessage

    - [x] SendMessageToPlayer

    - [x] GetVersion

    - [x] GetSkillLevel

    - [x] GetZAim

    - [x] GetSurfingOffsets

    - [x] GetRotationQuat

    - [x] GetPlayerSpectateID

    - [x] GetSpectateType

    - [x] GetRawIp

    - [x] SetGravity

    - [x] GetGravity

    - [x] SetAdmin

    - [x] IsSpawned

    - [x] IsControllable

    - [x] IsCameraTargetEnabled

    - [x] ToggleGhostMode

    - [x] GetGhostMode

    - [x] AllowWeapons

    - [x] AreWeaponsAllowed

    - [x] IsPlayerUsingOfficialClient

    - [x] GetAnimationFlags

    - [x] IsInDriveByMode

    - [x] IsCuffed

    - [x] IsUsingOmp

    - [x] IsInModShop

    - [x] GetSirenState

    - [x] GetLandingGearState

    - [x] GetHydraReactorAngle

    - [x] GetTrainSpeed

- [ ] **Actor**

    - [ ] Create

    - [ ] Destroy

    - [ ] FromID

    - [ ] GetID

    - [ ] IsStreamedInFor

    - [ ] SetVirtualWorld

    - [ ] GetVirtualWorld

    - [ ] ApplyAnimation

    - [ ] ClearAnimations

    - [ ] SetPos

    - [ ] GetPos

    - [ ] SetFacingAngle

    - [ ] GetFacingAngle

    - [ ] SetHealth

    - [ ] GetHealth

    - [ ] SetInvulnerable

    - [ ] IsInvulnerable

    - [ ] IsValid

    - [ ] SetSkin

    - [ ] GetSkin

    - [ ] GetAnimation

    - [ ] GetSpawnInfo

- [ ] **Checkpoint**

    - [ ] Set

    - [ ] Disable

    - [ ] IsPlayerIn

    - [ ] IsActive

    - [ ] Get

- [ ] **RaceCheckpoint**

    - [ ] Set

    - [ ] Disable

    - [ ] IsPlayerIn

    - [ ] IsActive

    - [ ] Get

- [ ] **Class**

    - [ ] Add

    - [ ] FromID

    - [ ] GetID

    - [ ] Count

    - [ ] GetData

    - [ ] Edit

- [ ] **Component**

    - [ ] Create

- [ ] **Config**

    - [ ] GetAsBool

    - [ ] GetAsInt

    - [ ] GetAsFloat

    - [ ] GetAsString

- [ ] **Core**

    - [ ] TickCount

    - [ ] MaxPlayers

    - [ ] Log

    - [ ] IsAdminTeleportAllowed

    - [ ] AllowAdminTeleport

    - [ ] AreAllAnimationsEnabled

    - [ ] EnableAllAnimations

    - [ ] IsAnimationLibraryValid

    - [ ] AreInteriorWeaponsAllowed

    - [ ] AllowInteriorWeapons

    - [ ] BlockIpAddress

    - [ ] UnBlockIpAddress

    - [ ] DisableEntryExitMarkers

    - [ ] DisableNameTagsLOS

    - [ ] EnableZoneNames

    - [ ] ShowGameTextForAll

    - [ ] HideGameTextForAll

    - [ ] NetworkStats

    - [ ] ServerTickRate

    - [ ] GetWeaponName

    - [ ] SetChatRadius

    - [ ] SetMarkerRadius

    - [ ] SendRconCommand

    - [ ] SetDeathDropAmount

    - [ ] GameMode_SetText

    - [ ] SetGravity

    - [ ] GetGravity

    - [ ] SetNameTagsDrawDistance

    - [ ] SetWeather

    - [ ] SetWorldTime

    - [ ] ShowNameTags

    - [ ] ShowPlayerMarkers

    - [ ] UsePedAnims

    - [ ] GetWeather

    - [ ] GetWorldTime

    - [ ] ToggleChatTextReplacement

    - [ ] IsChatTextReplacementToggled

    - [ ] IsNickNameValid

    - [ ] AllowNickNameCharacter

    - [ ] IsNickNameCharacterAllowed

    - [ ] ClearBanList

    - [ ] IsIpAddressBanned

    - [ ] GetWeaponSlot

    - [ ] AddRule

    - [ ] IsValidRule

    - [ ] RemoveRule

- [ ] **NPC**

    - [ ] Connect

    - [ ] Create

    - [ ] Destroy

    - [ ] FromID

    - [ ] GetID

    - [ ] IsValid

    - [ ] GetPlayer

    - [ ] Spawn

    - [ ] Respawn

    - [ ] SetPos

    - [ ] GetPos

    - [ ] SetRot

    - [ ] GetRot

    - [ ] SetFacingAngle

    - [ ] GetFacingAngle

    - [ ] SetVirtualWorld

    - [ ] GetVirtualWorld

    - [ ] SetInterior

    - [ ] GetInterior

    - [ ] Move

    - [ ] MoveToPlayer

    - [ ] StopMove

    - [ ] IsMoving

    - [ ] SetSkin

    - [ ] IsStreamedIn

    - [ ] IsAnyStreamedIn

    - [ ] GetAll

    - [ ] SetHealth

    - [ ] GetHealth

    - [ ] SetArmour

    - [ ] GetArmour

    - [ ] IsDead

    - [ ] SetInvulnerable

    - [ ] IsInvulnerable

    - [ ] SetWeapon

    - [ ] GetWeapon

    - [ ] SetAmmo

    - [ ] GetAmmo

    - [ ] SetAmmoInClip

    - [ ] GetAmmoInClip

    - [ ] EnableReloading

    - [ ] IsReloadEnabled

    - [ ] IsReloading

    - [ ] EnableInfiniteAmmo

    - [ ] IsInfiniteAmmoEnabled

    - [ ] GetWeaponState

    - [ ] SetKeys

    - [ ] GetKeys

    - [ ] SetWeaponSkillLevel

    - [ ] GetWeaponSkillLevel

    - [ ] MeleeAttack

    - [ ] StopMeleeAttack

    - [ ] IsMeleeAttacking

    - [ ] SetFightingStyle

    - [ ] GetFightingStyle

    - [ ] Shoot

    - [ ] IsShooting

    - [ ] AimAt

    - [ ] AimAtPlayer

    - [ ] StopAim

    - [ ] IsAiming

    - [ ] IsAimingAtPlayer

    - [ ] SetWeaponAccuracy

    - [ ] GetWeaponAccuracy

    - [ ] SetWeaponReloadTime

    - [ ] GetWeaponReloadTime

    - [ ] GetWeaponActualReloadTime

    - [ ] SetWeaponShootTime

    - [ ] GetWeaponShootTime

    - [ ] SetWeaponClipSize

    - [ ] GetWeaponClipSize

    - [ ] GetWeaponActualClipSize

    - [ ] EnterVehicle

    - [ ] ExitVehicle

    - [ ] PutInVehicle

    - [ ] RemoveFromVehicle

    - [ ] GetVehicle

    - [ ] GetVehicleID

    - [ ] GetEnteringVehicle

    - [ ] GetEnteringVehicleID

    - [ ] GetVehicleSeat

    - [ ] GetEnteringVehicleSeat

    - [ ] IsEnteringVehicle

    - [ ] UseVehicleSiren

    - [ ] IsVehicleSirenUsed

    - [ ] SetVehicleHealth

    - [ ] GetVehicleHealth

    - [ ] SetVehicleHydraThrusters

    - [ ] GetVehicleHydraThrusters

    - [ ] SetVehicleGearState

    - [ ] GetVehicleGearState

    - [ ] SetVehicleTrainSpeed

    - [ ] GetVehicleTrainSpeed

    - [ ] CreatePath

    - [ ] DestroyPath

    - [ ] DestroyAllPath

    - [ ] GetPathCount

    - [ ] AddPointToPath

    - [ ] RemovePointFromPath

    - [ ] ClearPath

    - [ ] GetPathPointCount

    - [ ] GetPathPoint

    - [ ] GetCurrentPathPointIndex

    - [ ] IsValidPath

    - [ ] HasPathPointInRange

    - [ ] MoveByPath

    - [ ] ResetAnimation

    - [ ] SetAnimation

    - [ ] GetAnimation

    - [ ] ApplyAnimation

    - [ ] ClearAnimations

    - [ ] SetSpecialAction

    - [ ] GetSpecialAction

    - [ ] StartPlayback

    - [ ] StartPlaybackEx

    - [ ] StopPlayback

    - [ ] PausePlayback

    - [ ] IsPlayingPlayback

    - [ ] IsPlaybackPaused

    - [ ] LoadRecord

    - [ ] UnloadRecord

    - [ ] IsValidRecord

    - [ ] GetRecordCount

    - [ ] UnloadAllRecords

    - [ ] OpenNode

    - [ ] CloseNode

    - [ ] IsNodeOpen

    - [ ] GetNodeType

    - [ ] SetNodePoint

    - [ ] GetNodePointPosition

    - [ ] GetNodePointCount

    - [ ] GetNodeInfo

    - [ ] PlayNode

    - [ ] StopPlayingNode

    - [ ] PausePlayingNode

    - [ ] ResumePlayingNode

    - [ ] IsPlayingNodePaused

    - [ ] IsPlayingNode

    - [ ] ChangeNode

    - [ ] UpdateNodePoint

    - [ ] SetSurfingOffset

    - [ ] GetSurfingOffset

    - [ ] SetSurfingVehicle

    - [ ] GetSurfingVehicle

    - [ ] SetSurfingObject

    - [ ] GetSurfingObject

    - [ ] SetSurfingPlayerObject

    - [ ] GetSurfingPlayerObject

    - [ ] ResetSurfingData

- [ ] **CustomModel**

    - [ ] AddCharModel

    - [ ] AddSimpleModel

    - [ ] AddSimpleModelTimed

    - [ ] RedirectDownload

    - [ ] FindModelFileNameFromCRC

    - [ ] IsValid

    - [ ] GetPath

- [ ] **Dialog**

    - [ ] Show

    - [ ] Hide

- [ ] **Event**

    - [ ] AddHandler

    - [ ] RemoveHandler

    - [ ] RemoveAllHandlers

- [ ] **GangZone**

    - [ ] Create

    - [ ] Destroy

    - [ ] FromID

    - [ ] GetID

    - [ ] ShowForPlayer

    - [ ] ShowForAll

    - [ ] HideForPlayer

    - [ ] HideForAll

    - [ ] FlashForPlayer

    - [ ] FlashForAll

    - [ ] StopFlashForPlayer

    - [ ] StopFlashForAll

    - [ ] IsValid

    - [ ] IsPlayerIn

    - [ ] IsVisibleForPlayer

    - [ ] GetColorForPlayer

    - [ ] GetFlashColorForPlayer

    - [ ] IsFlashingForPlayer

    - [ ] GetPos

    - [ ] UseCheck

- [ ] **Menu**

    - [ ] Create

    - [ ] Destroy

    - [ ] FromID

    - [ ] GetID

    - [ ] AddItem

    - [ ] SetColumnHeader

    - [ ] ShowForPlayer

    - [ ] HideForPlayer

    - [ ] Disable

    - [ ] DisableRow

    - [ ] IsValid

    - [ ] IsDisabled

    - [ ] IsRowDisabled

    - [ ] GetColumns

    - [ ] GetItems

    - [ ] GetPos

    - [ ] GetColumnWidth

    - [ ] GetColumnHeader

    - [ ] GetItem

- [ ] **Object**

    - [ ] Create

    - [ ] Destroy

    - [ ] FromID

    - [ ] GetID

    - [ ] AttachToVehicle

    - [ ] AttachToObject

    - [ ] AttachToPlayer

    - [ ] SetPos

    - [ ] GetPos

    - [ ] SetRot

    - [ ] GetRot

    - [ ] GetModel

    - [ ] SetNoCameraCollision

    - [ ] IsValid

    - [ ] Move

    - [ ] Stop

    - [ ] IsMoving

    - [ ] BeginEditing

    - [ ] BeginSelecting

    - [ ] EndEditing

    - [ ] SetMaterial

    - [ ] SetMaterialText

    - [ ] SetDefaultCameraCollision

    - [ ] GetDrawDistance

    - [ ] GetMoveSpeed

    - [ ] GetMovingTargetPos

    - [ ] GetMovingTargetRot

    - [ ] GetAttachedData

    - [ ] GetAttachedOffset

    - [ ] GetSyncRotation

    - [ ] IsMaterialSlotUsed

    - [ ] GetMaterial

    - [ ] GetMaterialText

    - [ ] IsObjectNoCameraCollision

    - [ ] GetType

- [ ] **PlayerObject**

    - [ ] Create

    - [ ] Destroy

    - [ ] FromID

    - [ ] GetID

    - [ ] AttachToVehicle

    - [ ] AttachToPlayer

    - [ ] AttachToObject

    - [ ] SetPos

    - [ ] GetPos

    - [ ] SetRot

    - [ ] GetRot

    - [ ] GetModel

    - [ ] SetNoCameraCollision

    - [ ] IsValid

    - [ ] Move

    - [ ] Stop

    - [ ] IsMoving

    - [ ] BeginEditing

    - [ ] SetMaterial

    - [ ] SetMaterialText

    - [ ] GetDrawDistance

    - [ ] GetMoveSpeed

    - [ ] GetMovingTargetPos

    - [ ] GetMovingTargetRot

    - [ ] GetAttachedData

    - [ ] GetAttachedOffset

    - [ ] GetSyncRotation

    - [ ] IsMaterialSlotUsed

    - [ ] GetMaterial

    - [ ] GetMaterialText

    - [ ] IsNoCameraCollision

- [ ] **Pickup**

    - [ ] Create

    - [ ] AddStatic

    - [ ] Destroy

    - [ ] FromID

    - [ ] GetID

    - [ ] IsValid

    - [ ] IsStreamedIn

    - [ ] GetPos

    - [ ] GetModel

    - [ ] GetType

    - [ ] GetVirtualWorld

    - [ ] SetPos

    - [ ] SetModel

    - [ ] SetType

    - [ ] SetVirtualWorld

    - [ ] ShowForPlayer

    - [ ] HideForPlayer

    - [ ] IsHiddenForPlayer

- [ ] **All**

    - [ ] SendClientMessage

    - [ ] CreateExplosion

    - [ ] SendDeathMessage

    - [ ] EnableStuntBonus

- [ ] **Recording**

    - [ ] Start

    - [ ] Stop

- [ ] **TextDraw**

    - [ ] Create

    - [ ] Destroy

    - [ ] FromID

    - [ ] GetID

    - [ ] IsValid

    - [ ] IsVisibleForPlayer

    - [ ] SetLetterSize

    - [ ] SetTextSize

    - [ ] SetAlignment

    - [ ] SetColor

    - [ ] SetUseBox

    - [ ] SetBoxColor

    - [ ] SetShadow

    - [ ] SetOutline

    - [ ] SetBackgroundColor

    - [ ] SetFont

    - [ ] SetProportional

    - [ ] SetSelectable

    - [ ] ShowForPlayer

    - [ ] HideForPlayer

    - [ ] ShowForAll

    - [ ] HideForAll

    - [ ] SetString

    - [ ] SetPreviewModel

    - [ ] SetPreviewRot

    - [ ] SetPreviewVehCol

    - [ ] SetPos

    - [ ] GetString

    - [ ] GetLetterSize

    - [ ] GetTextSize

    - [ ] GetPos

    - [ ] GetColor

    - [ ] GetBoxColor

    - [ ] GetBackgroundColor

    - [ ] GetShadow

    - [ ] GetOutline

    - [ ] GetFont

    - [ ] IsBox

    - [ ] IsProportional

    - [ ] IsSelectable

    - [ ] GetAlignment

    - [ ] GetPreviewModel

    - [ ] GetPreviewRot

    - [ ] GetPreviewVehColor

    - [ ] SetStringForPlayer

- [ ] **PlayerTextDraw**

    - [ ] Create

    - [ ] Destroy

    - [ ] FromID

    - [ ] GetID

    - [ ] IsValid

    - [ ] IsVisible

    - [ ] SetLetterSize

    - [ ] SetTextSize

    - [ ] SetAlignment

    - [ ] SetColor

    - [ ] UseBox

    - [ ] SetBoxColor

    - [ ] SetShadow

    - [ ] SetOutline

    - [ ] SetBackgroundColor

    - [ ] SetFont

    - [ ] SetProportional

    - [ ] SetSelectable

    - [ ] Show

    - [ ] Hide

    - [ ] SetString

    - [ ] SetPreviewModel

    - [ ] SetPreviewRot

    - [ ] SetPreviewVehCol

    - [ ] SetPos

    - [ ] GetString

    - [ ] GetLetterSize

    - [ ] GetTextSize

    - [ ] GetPos

    - [ ] GetColor

    - [ ] GetBoxColor

    - [ ] GetBackgroundColor

    - [ ] GetShadow

    - [ ] GetOutline

    - [ ] GetFont

    - [ ] IsBox

    - [ ] IsProportional

    - [ ] IsSelectable

    - [ ] GetAlignment

    - [ ] GetPreviewModel

    - [ ] GetPreviewRot

    - [ ] GetPreviewVehColor

- [ ] **TextLabel**

    - [ ] Create

    - [ ] Destroy

    - [ ] FromID

    - [ ] GetID

    - [ ] AttachToPlayer

    - [ ] AttachToVehicle

    - [ ] UpdateText

    - [ ] IsValid

    - [ ] IsStreamedIn

    - [ ] GetText

    - [ ] GetColor

    - [ ] GetPos

    - [ ] SetDrawDistance

    - [ ] GetDrawDistance

    - [ ] GetLOS

    - [ ] SetLOS

    - [ ] GetVirtualWorld

    - [ ] SetVirtualWorld

    - [ ] GetAttachedData

- [ ] **PlayerTextLabel**

    - [ ] Create

    - [ ] Destroy

    - [ ] FromID

    - [ ] GetID

    - [ ] UpdateText

    - [ ] IsValid

    - [ ] GetText

    - [ ] GetColor

    - [ ] GetPos

    - [ ] SetDrawDistance

    - [ ] GetDrawDistance

    - [ ] GetLOS

    - [ ] SetLOS

    - [ ] GetVirtualWorld

    - [ ] GetAttachedData

- [ ] **Vehicle**

    - [ ] Create

    - [ ] Destroy

    - [ ] FromID

    - [ ] GetID

    - [ ] GetMaxPassengerSeats

    - [ ] IsStreamedIn

    - [ ] GetPos

    - [ ] SetPos

    - [ ] GetZAngle

    - [ ] GetRotationQuat

    - [ ] GetDistanceFromPoint

    - [ ] SetZAngle

    - [ ] SetParamsForPlayer

    - [ ] UseManualEngineAndLights

    - [ ] SetParamsEx

    - [ ] GetParamsEx

    - [ ] GetParamsSirenState

    - [ ] SetParamsCarDoors

    - [ ] GetParamsCarDoors

    - [ ] SetParamsCarWindows

    - [ ] GetParamsCarWindows

    - [ ] SetToRespawn

    - [ ] LinkToInterior

    - [ ] AddComponent

    - [ ] RemoveComponent

    - [ ] ChangeColor

    - [ ] ChangePaintjob

    - [ ] SetHealth

    - [ ] GetHealth

    - [ ] AttachTrailer

    - [ ] DetachTrailer

    - [ ] IsTrailerAttached

    - [ ] GetTrailer

    - [ ] SetNumberPlate

    - [ ] GetModel

    - [ ] GetComponentInSlot

    - [ ] GetComponentType

    - [ ] CanHaveComponent

    - [ ] GetRandomColorPair

    - [ ] ColorIndexToColor

    - [ ] Repair

    - [ ] GetVelocity

    - [ ] SetVelocity

    - [ ] SetAngularVelocity

    - [ ] GetDamageStatus

    - [ ] UpdateDamageStatus

    - [ ] GetModelInfo

    - [ ] SetVirtualWorld

    - [ ] GetVirtualWorld

    - [ ] GetLandingGearState

    - [ ] IsValid

    - [ ] AddStatic

    - [ ] AddStaticEx

    - [ ] EnableFriendlyFire

    - [ ] GetSpawnInfo

    - [ ] SetSpawnInfo

    - [ ] GetModelCount

    - [ ] GetModelsUsed

    - [ ] GetPaintjob

    - [ ] GetColor

    - [ ] GetInterior

    - [ ] GetNumberPlate

    - [ ] SetRespawnDelay

    - [ ] GetRespawnDelay

    - [ ] GetCab

    - [ ] GetTower

    - [ ] GetOccupiedTick

    - [ ] GetRespawnTick

    - [ ] HasBeenOccupied

    - [ ] IsOccupied

    - [ ] IsDead

    - [ ] SetParamsSirenState

    - [ ] ToggleSirenEnabled

    - [ ] IsSirenEnabled

    - [ ] GetLastDriver

    - [ ] GetDriver

    - [ ] GetSirenState

    - [ ] GetHydraReactorAngle

    - [ ] GetTrainSpeed

    - [ ] GetMatrix

    - [ ] GetOccupant

    - [ ] CountOccupants

### Events

Implemented and tested events.

- [x] onPlayerConnect
- [x] onPlayerDisconnect
- [ ] onPlayerRequestClass
- [ ] onPlayerRequestSpawn
- [ ] onPlayerSpawn
- [ ] onPlayerDeath
- [ ] onPlayerUpdate
- [ ] onPlayerText
- [ ] onPlayerCommandText
- [ ] onPlayerInteriorChange
- [ ] onPlayerStateChange
- [ ] onPlayerKeyStateChange
- [ ] onPlayerEnterVehicle
- [ ] onPlayerExitVehicle
- [ ] onPlayerEnterCheckpoint
- [ ] onPlayerLeaveCheckpoint
- [ ] onPlayerEnterRaceCheckpoint
- [ ] onPlayerLeaveRaceCheckpoint
- [ ] onPlayerGiveDamage
- [ ] onPlayerTakeDamage
- [ ] onPlayerGiveDamageActor
- [ ] onPlayerWeaponShot
- [ ] onPlayerPickUpPickup
- [ ] onPlayerObjectMoved
- [ ] onPlayerEditObject
- [ ] onPlayerEditAttachedObject
- [ ] onPlayerSelectObject
- [ ] onPlayerClickMap
- [ ] onPlayerClickTextDraw
- [ ] onPlayerClickPlayerTextDraw
- [ ] onPlayerClickPlayer
- [ ] onPlayerStreamIn
- [ ] onPlayerStreamOut
- [ ] onPlayerExitedMenu
- [ ] onPlayerSelectedMenuRow
- [ ] onPlayerRequestDownload
- [ ] onTick
- [ ] onIncomingConnection
- [ ] onRconLoginAttempt
- [ ] onRconCommand
- [ ] onDialogResponse
- [ ] onVehicleSpawn
- [ ] onVehicleDeath
- [ ] onVehicleMod
- [ ] onVehiclePaintjob
- [ ] onVehicleRespray
- [ ] onVehicleDamageStatusUpdate
- [ ] onVehicleSirenStateChange
- [ ] onVehicleStreamIn
- [ ] onVehicleStreamOut
- [ ] onUnoccupiedVehicleUpdate
- [ ] onTrailerUpdate
- [ ] onActorStreamIn
- [ ] onActorStreamOut
- [ ] onObjectMoved
- [ ] onEnterExitModShop

Building command:

```sh
CGO_ENABLED=1 GOOS=linux GOARCH=386 go build -buildmode=c-shared -o basic.so examples/basic/main.go
```
