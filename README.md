# GOMP

Open Multiplayer C-API implementation for Golang.

Under development. Do not use in production under any circumstances.

C-API: https://github.com/openmultiplayer/open.mp-capi

### API

API wrappers.


- [ ] **Player**

    - [ ] SetSpawnInfo

    - [ ] GetSpawnInfo

    - [ ] GetNetworkStats

    - [ ] NetStatsBytesReceived

    - [ ] NetStatsBytesSent

    - [ ] NetStatsConnectionStatus

    - [ ] NetStatsGetConnectedTime

    - [ ] NetStatsGetIpPort

    - [ ] NetStatsMessagesReceived

    - [ ] NetStatsMessagesRecvPerSecond

    - [ ] NetStatsMessagesSent

    - [ ] NetStatsPacketLossPercent

    - [ ] GetCustomSkin

    - [ ] GetDialog

    - [ ] GetDialogData

    - [ ] GetMenu

    - [ ] GetSurfingPlayerObject

    - [ ] GetCameraTargetPlayerObject

    - [x] FromID

    - [x] GetID

    - [ ] SendClientMessage

    - [ ] SetCameraPos

    - [ ] SetDrunkLevel

    - [ ] SetInterior

    - [ ] SetWantedLevel

    - [ ] SetWeather

    - [ ] GetWeather

    - [ ] SetSkin

    - [ ] SetShopName

    - [ ] GiveMoney

    - [ ] SetCameraLookAt

    - [ ] SetCameraBehind

    - [ ] CreateExplosion

    - [ ] PlayAudioStream

    - [ ] StopAudioStream

    - [ ] ToggleWidescreen

    - [ ] IsWidescreenToggled

    - [ ] SetHealth

    - [ ] GetHealth

    - [ ] SetArmor

    - [ ] GetArmor

    - [ ] SetTeam

    - [ ] GetTeam

    - [ ] SetScore

    - [ ] GetScore

    - [ ] GetSkin

    - [ ] SetColor

    - [ ] GetColor

    - [ ] GetDefaultColor

    - [ ] GetDrunkLevel

    - [ ] GiveWeapon

    - [ ] RemoveWeapon

    - [ ] GetMoney

    - [ ] ResetMoney

    - [ ] SetName

    - [x] GetName

    - [ ] GetState

    - [ ] GetPing

    - [ ] GetWeapon

    - [ ] SetTime

    - [ ] GetTime

    - [ ] ToggleClock

    - [ ] HasClock

    - [ ] ForceClassSelection

    - [ ] GetWantedLevel

    - [ ] SetFightingStyle

    - [ ] GetFightingStyle

    - [ ] SetVelocity

    - [ ] GetVelocity

    - [ ] GetCameraPos

    - [ ] GetDistanceFromPoint

    - [ ] GetInterior

    - [ ] SetPos

    - [ ] GetPos

    - [ ] GetVirtualWorld

    - [ ] IsNPC

    - [ ] IsStreamedIn

    - [ ] PlayGameSound

    - [ ] SpectatePlayer

    - [ ] SpectateVehicle

    - [ ] SetVirtualWorld

    - [ ] SetWorldBounds

    - [ ] ClearWorldBounds

    - [ ] GetWorldBounds

    - [ ] ClearAnimations

    - [ ] GetLastShotVectors

    - [ ] GetCameraTargetPlayer

    - [ ] GetCameraTargetActor

    - [ ] GetCameraTargetObject

    - [ ] GetCameraTargetVehicle

    - [ ] PutInVehicle

    - [ ] RemoveBuilding

    - [ ] GetBuildingsRemoved

    - [ ] RemoveFromVehicle

    - [ ] RemoveMapIcon

    - [ ] SetMapIcon

    - [ ] ResetWeapons

    - [ ] SetAmmo

    - [ ] SetArmedWeapon

    - [ ] SetChatBubble

    - [ ] SetPosFindZ

    - [ ] SetSkillLevel

    - [ ] SetSpecialAction

    - [ ] ShowNameTagForPlayer

    - [ ] ToggleControllable

    - [ ] ToggleSpectating

    - [ ] ApplyAnimation

    - [ ] GetAnimationName

    - [ ] EditAttachedObject

    - [ ] EnableCameraTarget

    - [ ] EnableStuntBonus

    - [ ] GetPlayerAmmo

    - [ ] GetAnimationIndex

    - [ ] GetFacingAngle

    - [ ] GetIp

    - [ ] GetSpecialAction

    - [ ] GetVehicleID

    - [ ] GetVehicleSeat

    - [ ] GetWeaponData

    - [ ] GetWeaponState

    - [ ] InterpolateCameraPos

    - [ ] InterpolateCameraLookAt

    - [ ] IsPlayerAttachedObjectSlotUsed

    - [ ] AttachCameraToObject

    - [ ] AttachCameraToPlayerObject

    - [ ] GetCameraAspectRatio

    - [ ] GetCameraFrontVector

    - [ ] GetCameraMode

    - [ ] GetKeys

    - [ ] GetSurfingVehicle

    - [ ] GetSurfingObject

    - [ ] GetTargetPlayer

    - [ ] GetTargetActor

    - [ ] IsInVehicle

    - [ ] IsInAnyVehicle

    - [ ] IsInRangeOfPoint

    - [ ] PlayCrimeReport

    - [ ] RemoveAttachedObject

    - [ ] SetAttachedObject

    - [ ] GetAttachedObject

    - [ ] SetFacingAngle

    - [ ] SetMarkerForPlayer

    - [ ] GetMarkerForPlayer

    - [ ] AllowTeleport

    - [ ] IsTeleportAllowed

    - [ ] DisableRemoteVehicleCollisions

    - [ ] GetCameraZoom

    - [ ] SelectTextDraw

    - [ ] CancelSelectTextDraw

    - [ ] SendClientCheck

    - [ ] Spawn

    - [ ] GPCI

    - [ ] IsAdmin

    - [ ] Kick

    - [ ] ShowGameText

    - [ ] HideGameText

    - [ ] HasGameText

    - [ ] GetGameText

    - [ ] Ban

    - [ ] BanEx

    - [ ] SendDeathMessage

    - [ ] SendMessageToPlayer

    - [ ] GetVersion

    - [ ] GetSkillLevel

    - [ ] GetZAim

    - [ ] GetSurfingOffsets

    - [ ] GetRotationQuat

    - [ ] GetPlayerSpectateID

    - [ ] GetSpectateType

    - [ ] GetRawIp

    - [ ] SetGravity

    - [ ] GetGravity

    - [ ] SetAdmin

    - [ ] IsSpawned

    - [ ] IsControllable

    - [ ] IsCameraTargetEnabled

    - [ ] ToggleGhostMode

    - [ ] GetGhostMode

    - [ ] AllowWeapons

    - [ ] AreWeaponsAllowed

    - [ ] IsPlayerUsingOfficialClient

    - [ ] GetAnimationFlags

    - [ ] IsInDriveByMode

    - [ ] IsCuffed

    - [ ] IsUsingOmp

    - [ ] IsInModShop

    - [ ] GetSirenState

    - [ ] GetLandingGearState

    - [ ] GetHydraReactorAngle

    - [ ] GetTrainSpeed

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
