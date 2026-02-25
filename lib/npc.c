#include "npc.h"
#include "main.h"

bool NPC_Connect(const char *name, const char *script) {
  return api.NPC.Connect(name, script);
}

void *NPC_Create(const char *name, int *id) {
  return api.NPC.Create(name, id);
}

bool NPC_Destroy(void *npc) {
  return api.NPC.Destroy(npc);
}

void *NPC_FromID(int npcid) {
  return api.NPC.FromID(npcid);
}

int NPC_GetID(void *npc) {
  return api.NPC.GetID(npc);
}

bool NPC_IsValid(void *npc) {
  return api.NPC.IsValid(npc);
}

void *NPC_GetPlayer(void *npc) {
  return api.NPC.GetPlayer(npc);
}

bool NPC_Spawn(void *npc) {
  return api.NPC.Spawn(npc);
}

bool NPC_Respawn(void *npc) {
  return api.NPC.Respawn(npc);
}

bool NPC_SetPos(void *npc, float x, float y, float z) {
  return api.NPC.SetPos(npc, x, y, z);
}

bool NPC_GetPos(void *npc, float *x, float *y, float *z) {
  return api.NPC.GetPos(npc, x, y, z);
}

bool NPC_SetRot(void *npc, float rx, float ry, float rz) {
  return api.NPC.SetRot(npc, rx, ry, rz);
}

bool NPC_GetRot(void *npc, float *rx, float *ry, float *rz) {
  return api.NPC.GetRot(npc, rx, ry, rz);
}

bool NPC_SetFacingAngle(void *npc, float angle) {
  return api.NPC.SetFacingAngle(npc, angle);
}

bool NPC_GetFacingAngle(void *npc, float *angle) {
  return api.NPC.GetFacingAngle(npc, angle);
}

bool NPC_SetVirtualWorld(void *npc, int virtualWorld) {
  return api.NPC.SetVirtualWorld(npc, virtualWorld);
}

int NPC_GetVirtualWorld(void *npc) {
  return api.NPC.GetVirtualWorld(npc);
}

bool NPC_SetInterior(void *npc, int interior) {
  return api.NPC.SetInterior(npc, interior);
}

int NPC_GetInterior(void *npc) {
  return api.NPC.GetInterior(npc);
}

bool NPC_Move(void *npc, float x, float y, float z, int moveType,
              float moveSpeed, float stopRange) {
  return api.NPC.Move(npc, x, y, z, moveType, moveSpeed, stopRange);
}

bool NPC_MoveToPlayer(void *npc, void *player, int moveType, float moveSpeed,
                      float stopRange, int posCheckUpdateDelay,
                      bool autoRestart) {
  return api.NPC.MoveToPlayer(npc, player, moveType, moveSpeed, stopRange,
                               posCheckUpdateDelay, autoRestart);
}

bool NPC_StopMove(void *npc) {
  return api.NPC.StopMove(npc);
}

bool NPC_IsMoving(void *npc) {
  return api.NPC.IsMoving(npc);
}

bool NPC_SetSkin(void *npc, int model) {
  return api.NPC.SetSkin(npc, model);
}

bool NPC_IsStreamedIn(void *npc, void *player) {
  return api.NPC.IsStreamedIn(npc, player);
}

bool NPC_IsAnyStreamedIn(void *npc) {
  return api.NPC.IsAnyStreamedIn(npc);
}

int NPC_GetAll(int *npcsArr, int maxNPCs) {
  return api.NPC.GetAll(npcsArr, maxNPCs);
}

bool NPC_SetHealth(void *npc, float health) {
  return api.NPC.SetHealth(npc, health);
}

float NPC_GetHealth(void *npc) {
  return api.NPC.GetHealth(npc);
}

bool NPC_SetArmour(void *npc, float armour) {
  return api.NPC.SetArmour(npc, armour);
}

float NPC_GetArmour(void *npc) {
  return api.NPC.GetArmour(npc);
}

bool NPC_IsDead(void *npc) {
  return api.NPC.IsDead(npc);
}

bool NPC_SetInvulnerable(void *npc, bool toggle) {
  return api.NPC.SetInvulnerable(npc, toggle);
}

bool NPC_IsInvulnerable(void *npc) {
  return api.NPC.IsInvulnerable(npc);
}

bool NPC_SetWeapon(void *npc, uint8_t weapon) {
  return api.NPC.SetWeapon(npc, weapon);
}

uint8_t NPC_GetWeapon(void *npc) {
  return api.NPC.GetWeapon(npc);
}

bool NPC_SetAmmo(void *npc, int ammo) {
  return api.NPC.SetAmmo(npc, ammo);
}

int NPC_GetAmmo(void *npc) {
  return api.NPC.GetAmmo(npc);
}

bool NPC_SetAmmoInClip(void *npc, int ammo) {
  return api.NPC.SetAmmoInClip(npc, ammo);
}

int NPC_GetAmmoInClip(void *npc) {
  return api.NPC.GetAmmoInClip(npc);
}

bool NPC_EnableReloading(void *npc, bool enable) {
  return api.NPC.EnableReloading(npc, enable);
}

bool NPC_IsReloadEnabled(void *npc) {
  return api.NPC.IsReloadEnabled(npc);
}

bool NPC_IsReloading(void *npc) {
  return api.NPC.IsReloading(npc);
}

bool NPC_EnableInfiniteAmmo(void *npc, bool enable) {
  return api.NPC.EnableInfiniteAmmo(npc, enable);
}

bool NPC_IsInfiniteAmmoEnabled(void *npc) {
  return api.NPC.IsInfiniteAmmoEnabled(npc);
}

int NPC_GetWeaponState(void *npc) {
  return api.NPC.GetWeaponState(npc);
}

bool NPC_SetKeys(void *npc, uint16_t upAndDown, uint16_t leftAndRight,
                 uint16_t keys) {
  return api.NPC.SetKeys(npc, upAndDown, leftAndRight, keys);
}

bool NPC_GetKeys(void *npc, uint16_t *upAndDown, uint16_t *leftAndRight,
                 uint16_t *keys) {
  return api.NPC.GetKeys(npc, upAndDown, leftAndRight, keys);
}

bool NPC_SetWeaponSkillLevel(void *npc, uint8_t skill, int level) {
  return api.NPC.SetWeaponSkillLevel(npc, skill, level);
}

int NPC_GetWeaponSkillLevel(void *npc, int skill) {
  return api.NPC.GetWeaponSkillLevel(npc, skill);
}

bool NPC_MeleeAttack(void *npc, int time, bool secondaryAttack) {
  return api.NPC.MeleeAttack(npc, time, secondaryAttack);
}

bool NPC_StopMeleeAttack(void *npc) {
  return api.NPC.StopMeleeAttack(npc);
}

bool NPC_IsMeleeAttacking(void *npc) {
  return api.NPC.IsMeleeAttacking(npc);
}

bool NPC_SetFightingStyle(void *npc, int style) {
  return api.NPC.SetFightingStyle(npc, style);
}

int NPC_GetFightingStyle(void *npc) {
  return api.NPC.GetFightingStyle(npc);
}

bool NPC_Shoot(void *npc, uint8_t weapon, int hitId, int hitType, float endX,
               float endY, float endZ, float offsetX, float offsetY,
               float offsetZ, bool isHit, uint8_t checkInBetweenFlags) {
  return api.NPC.Shoot(npc, weapon, hitId, hitType, endX, endY, endZ, offsetX,
                       offsetY, offsetZ, isHit, checkInBetweenFlags);
}

bool NPC_IsShooting(void *npc) {
  return api.NPC.IsShooting(npc);
}

bool NPC_AimAt(void *npc, float x, float y, float z, bool shoot,
               int shootDelay, bool updateAngle, float offsetFromX,
               float offsetFromY, float offsetFromZ,
               uint8_t checkInBetweenFlags) {
  return api.NPC.AimAt(npc, x, y, z, shoot, shootDelay, updateAngle,
                       offsetFromX, offsetFromY, offsetFromZ,
                       checkInBetweenFlags);
}

bool NPC_AimAtPlayer(void *npc, void *atPlayer, bool shoot, int shootDelay,
                     bool updateAngle, float offsetX, float offsetY,
                     float offsetZ, float offsetFromX, float offsetFromY,
                     float offsetFromZ, uint8_t checkInBetweenFlags) {
  return api.NPC.AimAtPlayer(npc, atPlayer, shoot, shootDelay, updateAngle,
                              offsetX, offsetY, offsetZ, offsetFromX,
                              offsetFromY, offsetFromZ, checkInBetweenFlags);
}

bool NPC_StopAim(void *npc) {
  return api.NPC.StopAim(npc);
}

bool NPC_IsAiming(void *npc) {
  return api.NPC.IsAiming(npc);
}

bool NPC_IsAimingAtPlayer(void *npc, void *atPlayer) {
  return api.NPC.IsAimingAtPlayer(npc, atPlayer);
}

bool NPC_SetWeaponAccuracy(void *npc, int weapon, float accuracy) {
  return api.NPC.SetWeaponAccuracy(npc, weapon, accuracy);
}

float NPC_GetWeaponAccuracy(void *npc, int weapon) {
  return api.NPC.GetWeaponAccuracy(npc, weapon);
}

bool NPC_SetWeaponReloadTime(void *npc, int weapon, int time) {
  return api.NPC.SetWeaponReloadTime(npc, weapon, time);
}

int NPC_GetWeaponReloadTime(void *npc, int weapon) {
  return api.NPC.GetWeaponReloadTime(npc, weapon);
}

int NPC_GetWeaponActualReloadTime(void *npc, int weapon) {
  return api.NPC.GetWeaponActualReloadTime(npc, weapon);
}

bool NPC_SetWeaponShootTime(void *npc, int weapon, int time) {
  return api.NPC.SetWeaponShootTime(npc, weapon, time);
}

int NPC_GetWeaponShootTime(void *npc, int weapon) {
  return api.NPC.GetWeaponShootTime(npc, weapon);
}

bool NPC_SetWeaponClipSize(void *npc, int weapon, int size) {
  return api.NPC.SetWeaponClipSize(npc, weapon, size);
}

int NPC_GetWeaponClipSize(void *npc, int weapon) {
  return api.NPC.GetWeaponClipSize(npc, weapon);
}

int NPC_GetWeaponActualClipSize(void *npc, int weapon) {
  return api.NPC.GetWeaponActualClipSize(npc, weapon);
}

bool NPC_EnterVehicle(void *npc, void *vehicle, int seatId, int moveType) {
  return api.NPC.EnterVehicle(npc, vehicle, seatId, moveType);
}

bool NPC_ExitVehicle(void *npc) {
  return api.NPC.ExitVehicle(npc);
}

bool NPC_PutInVehicle(void *npc, void *vehicle, int seatId) {
  return api.NPC.PutInVehicle(npc, vehicle, seatId);
}

bool NPC_RemoveFromVehicle(void *npc) {
  return api.NPC.RemoveFromVehicle(npc);
}

void *NPC_GetVehicle(void *npc) {
  return api.NPC.GetVehicle(npc);
}

int NPC_GetVehicleID(void *npc) {
  return api.NPC.GetVehicleID(npc);
}

void *NPC_GetEnteringVehicle(void *npc) {
  return api.NPC.GetEnteringVehicle(npc);
}

int NPC_GetEnteringVehicleID(void *npc) {
  return api.NPC.GetEnteringVehicleID(npc);
}

int NPC_GetVehicleSeat(void *npc) {
  return api.NPC.GetVehicleSeat(npc);
}

int NPC_GetEnteringVehicleSeat(void *npc) {
  return api.NPC.GetEnteringVehicleSeat(npc);
}

bool NPC_IsEnteringVehicle(void *npc) {
  return api.NPC.IsEnteringVehicle(npc);
}

bool NPC_UseVehicleSiren(void *npc, bool use) {
  return api.NPC.UseVehicleSiren(npc, use);
}

bool NPC_IsVehicleSirenUsed(void *npc) {
  return api.NPC.IsVehicleSirenUsed(npc);
}

bool NPC_SetVehicleHealth(void *npc, float health) {
  return api.NPC.SetVehicleHealth(npc, health);
}

float NPC_GetVehicleHealth(void *npc) {
  return api.NPC.GetVehicleHealth(npc);
}

bool NPC_SetVehicleHydraThrusters(void *npc, int direction) {
  return api.NPC.SetVehicleHydraThrusters(npc, direction);
}

int NPC_GetVehicleHydraThrusters(void *npc) {
  return api.NPC.GetVehicleHydraThrusters(npc);
}

bool NPC_SetVehicleGearState(void *npc, int gearState) {
  return api.NPC.SetVehicleGearState(npc, gearState);
}

int NPC_GetVehicleGearState(void *npc) {
  return api.NPC.GetVehicleGearState(npc);
}

bool NPC_SetVehicleTrainSpeed(void *npc, float speed) {
  return api.NPC.SetVehicleTrainSpeed(npc, speed);
}

float NPC_GetVehicleTrainSpeed(void *npc) {
  return api.NPC.GetVehicleTrainSpeed(npc);
}

int NPC_CreatePath() {
  return api.NPC.CreatePath();
}

bool NPC_DestroyPath(int pathId) {
  return api.NPC.DestroyPath(pathId);
}

bool NPC_DestroyAllPath() {
  return api.NPC.DestroyAllPath();
}

int NPC_GetPathCount() {
  return api.NPC.GetPathCount();
}

bool NPC_AddPointToPath(int pathId, float x, float y, float z, float stopRange) {
  return api.NPC.AddPointToPath(pathId, x, y, z, stopRange);
}

bool NPC_RemovePointFromPath(int pathId, int pointIndex) {
  return api.NPC.RemovePointFromPath(pathId, pointIndex);
}

bool NPC_ClearPath(int pathId) {
  return api.NPC.ClearPath(pathId);
}

int NPC_GetPathPointCount(int pathId) {
  return api.NPC.GetPathPointCount(pathId);
}

bool NPC_GetPathPoint(int pathId, int pointIndex, float *x, float *y, float *z,
                      float *stopRange) {
  return api.NPC.GetPathPoint(pathId, pointIndex, x, y, z, stopRange);
}

int NPC_GetCurrentPathPointIndex(void *npc) {
  return api.NPC.GetCurrentPathPointIndex(npc);
}

bool NPC_IsValidPath(int pathId) {
  return api.NPC.IsValidPath(pathId);
}

bool NPC_HasPathPointInRange(int pathId, float x, float y, float z,
                              float radius) {
  return api.NPC.HasPathPointInRange(pathId, x, y, z, radius);
}

bool NPC_MoveByPath(void *npc, int pathId, int moveType, float moveSpeed,
                    bool reverse) {
  return api.NPC.MoveByPath(npc, pathId, moveType, moveSpeed, reverse);
}

bool NPC_ResetAnimation(void *npc) {
  return api.NPC.ResetAnimation(npc);
}

bool NPC_SetAnimation(void *npc, int animationId, float delta, bool loop,
                      bool lockX, bool lockY, bool freeze, int time) {
  return api.NPC.SetAnimation(npc, animationId, delta, loop, lockX, lockY,
                               freeze, time);
}

bool NPC_GetAnimation(void *npc, int *animationId, float *delta, bool *loop,
                      bool *lockX, bool *lockY, bool *freeze, int *time) {
  return api.NPC.GetAnimation(npc, animationId, delta, loop, lockX, lockY,
                               freeze, time);
}

bool NPC_ApplyAnimation(void *npc, const char *animlib, const char *animname,
                        float delta, bool loop, bool lockX, bool lockY,
                        bool freeze, int time) {
  return api.NPC.ApplyAnimation(npc, animlib, animname, delta, loop, lockX,
                                 lockY, freeze, time);
}

bool NPC_ClearAnimations(void *npc) {
  return api.NPC.ClearAnimations(npc);
}

bool NPC_SetSpecialAction(void *npc, int action) {
  return api.NPC.SetSpecialAction(npc, action);
}

int NPC_GetSpecialAction(void *npc) {
  return api.NPC.GetSpecialAction(npc);
}

bool NPC_StartPlayback(void *npc, const char *recordName, bool autoUnload,
                       float startPosX, float startPosY, float startPosZ,
                       float startRotX, float startRotY, float startRotZ) {
  return api.NPC.StartPlayback(npc, recordName, autoUnload, startPosX,
                                startPosY, startPosZ, startRotX, startRotY,
                                startRotZ);
}

bool NPC_StartPlaybackEx(void *npc, int recordId, bool autoUnload,
                         float startPosX, float startPosY, float startPosZ,
                         float startRotX, float startRotY, float startRotZ) {
  return api.NPC.StartPlaybackEx(npc, recordId, autoUnload, startPosX,
                                  startPosY, startPosZ, startRotX, startRotY,
                                  startRotZ);
}

bool NPC_StopPlayback(void *npc) {
  return api.NPC.StopPlayback(npc);
}

bool NPC_PausePlayback(void *npc, bool paused) {
  return api.NPC.PausePlayback(npc, paused);
}

bool NPC_IsPlayingPlayback(void *npc) {
  return api.NPC.IsPlayingPlayback(npc);
}

bool NPC_IsPlaybackPaused(void *npc) {
  return api.NPC.IsPlaybackPaused(npc);
}

int NPC_LoadRecord(const char *filePath) {
  return api.NPC.LoadRecord(filePath);
}

bool NPC_UnloadRecord(int recordId) {
  return api.NPC.UnloadRecord(recordId);
}

bool NPC_IsValidRecord(int recordId) {
  return api.NPC.IsValidRecord(recordId);
}

int NPC_GetRecordCount() {
  return api.NPC.GetRecordCount();
}

bool NPC_UnloadAllRecords() {
  return api.NPC.UnloadAllRecords();
}

bool NPC_OpenNode(int nodeId) {
  return api.NPC.OpenNode(nodeId);
}

bool NPC_CloseNode(int nodeId) {
  return api.NPC.CloseNode(nodeId);
}

bool NPC_IsNodeOpen(int nodeId) {
  return api.NPC.IsNodeOpen(nodeId);
}

int NPC_GetNodeType(int nodeId) {
  return api.NPC.GetNodeType(nodeId);
}

bool NPC_SetNodePoint(int nodeId, int pointId) {
  return api.NPC.SetNodePoint(nodeId, pointId);
}

bool NPC_GetNodePointPosition(int nodeId, float *x, float *y, float *z) {
  return api.NPC.GetNodePointPosition(nodeId, x, y, z);
}

int NPC_GetNodePointCount(int nodeId) {
  return api.NPC.GetNodePointCount(nodeId);
}

bool NPC_GetNodeInfo(int nodeId, uint32_t *vehicleNodes, uint32_t *pedNodes,
                     uint32_t *naviNodes) {
  return api.NPC.GetNodeInfo(nodeId, vehicleNodes, pedNodes, naviNodes);
}

bool NPC_PlayNode(void *npc, int nodeId, int moveType, float moveSpeed,
                  float radius, bool setAngle) {
  return api.NPC.PlayNode(npc, nodeId, moveType, moveSpeed, radius, setAngle);
}

bool NPC_StopPlayingNode(void *npc) {
  return api.NPC.StopPlayingNode(npc);
}

bool NPC_PausePlayingNode(void *npc) {
  return api.NPC.PausePlayingNode(npc);
}

bool NPC_ResumePlayingNode(void *npc) {
  return api.NPC.ResumePlayingNode(npc);
}

bool NPC_IsPlayingNodePaused(void *npc) {
  return api.NPC.IsPlayingNodePaused(npc);
}

bool NPC_IsPlayingNode(void *npc) {
  return api.NPC.IsPlayingNode(npc);
}

int NPC_ChangeNode(void *npc, int nodeId, int linkId) {
  return api.NPC.ChangeNode(npc, nodeId, linkId);
}

bool NPC_UpdateNodePoint(void *npc, int pointId) {
  return api.NPC.UpdateNodePoint(npc, pointId);
}

bool NPC_SetSurfingOffset(void *npc, float x, float y, float z) {
  return api.NPC.SetSurfingOffset(npc, x, y, z);
}

bool NPC_GetSurfingOffset(void *npc, float *x, float *y, float *z) {
  return api.NPC.GetSurfingOffset(npc, x, y, z);
}

bool NPC_SetSurfingVehicle(void *npc, void *vehicle) {
  return api.NPC.SetSurfingVehicle(npc, vehicle);
}

int NPC_GetSurfingVehicle(void *npc) {
  return api.NPC.GetSurfingVehicle(npc);
}

bool NPC_SetSurfingObject(void *npc, void *object) {
  return api.NPC.SetSurfingObject(npc, object);
}

int NPC_GetSurfingObject(void *npc) {
  return api.NPC.GetSurfingObject(npc);
}

bool NPC_SetSurfingPlayerObject(void *npc, void *player, int objectId) {
  return api.NPC.SetSurfingPlayerObject(npc, player, objectId);
}

int NPC_GetSurfingPlayerObject(void *npc) {
  return api.NPC.GetSurfingPlayerObject(npc);
}

bool NPC_ResetSurfingData(void *npc) {
  return api.NPC.ResetSurfingData(npc);
}
