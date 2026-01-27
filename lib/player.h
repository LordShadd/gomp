#ifndef PLAYER_H
#define PLAYER_H

#include "ompcapi.h"
#include <stdbool.h>
#include <stdint.h>

int Player_NetStatsBytesReceived(void *player);
int Player_NetStatsBytesSent(void *player);
int Player_NetStatsConnectionStatus(void *player);
int Player_NetStatsGetConnectedTime(void *player);
bool Player_NetStatsGetIpPort(void *player, struct CAPIStringBuffer *output);
int Player_NetStatsMessagesReceived(void *player);
int Player_NetStatsMessagesRecvPerSecond(void *player);
int Player_NetStatsMessagesSent(void *player);
float Player_NetStatsPacketLossPercent(void *player);
int Player_GetCustomSkin(void *player);
int Player_GetDialog(void *player);
bool Player_GetDialogData(void *player, int *dialogid, int *style,
                          struct CAPIStringView *title,
                          struct CAPIStringView *body,
                          struct CAPIStringView *button1,
                          struct CAPIStringView *button2);
void *Player_GetMenu(void *player);
void *Player_GetSurfingPlayerObject(void *player);
void *Player_GetCameraTargetPlayerObject(void *player);
void *Player_FromID(int playerid);
int Player_GetID(void *player);
bool Player_SendClientMessage(void *player, uint32_t color, const char *text);
bool Player_SetCameraPos(void *player, float x, float y, float z);
bool Player_SetDrunkLevel(void *player, int level);
bool Player_SetInterior(void *player, int interior);
bool Player_SetWantedLevel(void *player, int level);
bool Player_SetWeather(void *player, int weather);
int Player_GetWeather(void *player);
bool Player_SetSkin(void *player, int skin);
bool Player_SetShopName(void *player, const char *name);
bool Player_GiveMoney(void *player, int amount);
bool Player_SetCameraLookAt(void *player, float x, float y, float z,
                            int cutType);
bool Player_SetCameraBehind(void *player);
bool Player_CreateExplosion(void *player, float x, float y, float z, int type,
                            float radius);
bool Player_PlayAudioStream(void *player, const char *url, float x, float y,
                            float z, float distance, bool usePos);
bool Player_StopAudioStream(void *player);
bool Player_ToggleWidescreen(void *player, bool enable);
bool Player_IsWidescreenToggled(void *player);
bool Player_SetHealth(void *player, float health);
float Player_GetHealth(void *player);
bool Player_SetArmor(void *player, float armor);
float Player_GetArmor(void *player);
bool Player_SetTeam(void *player, int team);
int Player_GetTeam(void *player);
bool Player_SetScore(void *player, int score);
int Player_GetScore(void *player);
int Player_GetSkin(void *player);
bool Player_SetColor(void *player, uint32_t color);
uint32_t Player_GetColor(void *player);
uint32_t Player_GetDefaultColor(void *player);
int Player_GetDrunkLevel(void *player);
bool Player_GiveWeapon(void *player, int weapon, int ammo);
bool Player_RemoveWeapon(void *player, int weapon);
int Player_GetMoney(void *player);
bool Player_ResetMoney(void *player);
int Player_SetName(void *player, const char *name);
int Player_GetName(void *player, struct CAPIStringView *name);
int Player_GetState(void *player);
int Player_GetPing(void *player);
int Player_GetWeapon(void *player);
bool Player_SetTime(void *player, int hour, int minute);
bool Player_GetTime(void *player, int *hour, int *minute);
bool Player_ToggleClock(void *player, bool enable);
bool Player_HasClock(void *player);
bool Player_ForceClassSelection(void *player);
int Player_GetWantedLevel(void *player);
bool Player_SetFightingStyle(void *player, int style);
int Player_GetFightingStyle(void *player);
bool Player_SetVelocity(void *player, float x, float y, float z);
bool Player_GetVelocity(void *player, float *x, float *y, float *z);
bool Player_GetCameraPos(void *player, float *x, float *y, float *z);
float Player_GetDistanceFromPoint(void *player, float x, float y, float z);
int Player_GetInterior(void *player);
bool Player_SetPos(void *player, float x, float y, float z);
bool Player_GetPos(void *player, float *x, float *y, float *z);
int Player_GetVirtualWorld(void *player);
bool Player_IsNPC(void *player);
bool Player_IsStreamedIn(void *player, void *other);
bool Player_PlayGameSound(void *player, int sound, float x, float y, float z);
bool Player_SpectatePlayer(void *player, void *target, int mode);
bool Player_SpectateVehicle(void *player, void *target, int mode);
bool Player_SetVirtualWorld(void *player, int vw);
bool Player_SetWorldBounds(void *player, float xMax, float xMin, float yMax,
                           float yMin);
bool Player_ClearWorldBounds(void *player);
bool Player_GetWorldBounds(void *player, float *xmax, float *xmin, float *ymax,
                           float *ymin);
bool Player_ClearAnimations(void *player, int syncType);
bool Player_GetLastShotVectors(void *player, float *origin_x, float *origin_y,
                               float *origin_z, float *hit_x, float *hit_y,
                               float *hit_z);
void *Player_GetCameraTargetPlayer(void *player);
void *Player_GetCameraTargetActor(void *player);
void *Player_GetCameraTargetObject(void *player);
void *Player_GetCameraTargetVehicle(void *player);
bool Player_PutInVehicle(void *player, void *vehicle, int seat);
bool Player_RemoveBuilding(void *player, int model, float x, float y, float z,
                           float radius);
int Player_GetBuildingsRemoved(void *player);
bool Player_RemoveFromVehicle(void *player, bool force);
bool Player_RemoveMapIcon(void *player, int icon);
bool Player_SetMapIcon(void *player, int iconID, float x, float y, float z,
                       int type, uint32_t color, int style);
bool Player_ResetWeapons(void *player);
bool Player_SetAmmo(void *player, uint8_t id, uint32_t ammo);
bool Player_SetArmedWeapon(void *player, uint8_t weapon);
bool Player_SetChatBubble(void *player, const char *text, uint32_t color,
                          float drawdistance, int expiretime);
bool Player_SetPosFindZ(void *player, float x, float y, float z);
bool Player_SetSkillLevel(void *player, uint8_t weapon, int level);
bool Player_SetSpecialAction(void *player, uint32_t action);
bool Player_ShowNameTagForPlayer(void *player, void *other, bool enable);
bool Player_ToggleControllable(void *player, bool enable);
bool Player_ToggleSpectating(void *player, bool enable);
bool Player_ApplyAnimation(void *player, const char *animlib,
                           const char *animname, float delta, bool loop,
                           bool lockX, bool lockY, bool freeze, uint32_t time,
                           int sync);
bool Player_GetAnimationName(int index, struct CAPIStringView *lib,
                             struct CAPIStringView *name);
bool Player_EditAttachedObject(void *player, int index);
bool Player_EnableCameraTarget(void *player, bool enable);
bool Player_EnableStuntBonus(void *player, bool enable);
int Player_GetPlayerAmmo(void *player);
int Player_GetAnimationIndex(void *player);
float Player_GetFacingAngle(void *player);
int Player_GetIp(void *player, struct CAPIStringBuffer *ip);
int Player_GetSpecialAction(void *player);
int Player_GetVehicleID(void *player);
int Player_GetVehicleSeat(void *player);
bool Player_GetWeaponData(void *player, int slot, int *weaponid, int *ammo);
int Player_GetWeaponState(void *player);
bool Player_InterpolateCameraPos(void *player, float from_x, float from_y,
                                 float from_z, float to_x, float to_y,
                                 float to_z, int time, int cut);
bool Player_InterpolateCameraLookAt(void *player, float from_x, float from_y,
                                    float from_z, float to_x, float to_y,
                                    float to_z, int time, int cut);
bool Player_IsPlayerAttachedObjectSlotUsed(void *player, int index);
bool Player_AttachCameraToObject(void *player, void *object);
bool Player_AttachCameraToPlayerObject(void *player, void *object);
float Player_GetCameraAspectRatio(void *player);
bool Player_GetCameraFrontVector(void *player, float *x, float *y, float *z);
int Player_GetCameraMode(void *player);
bool Player_GetKeys(void *player, int *keys, int *updown, int *leftright);
void *Player_GetSurfingVehicle(void *player);
void *Player_GetSurfingObject(void *player);
void *Player_GetTargetPlayer(void *player);
void *Player_GetTargetActor(void *player);
bool Player_IsInVehicle(void *player, void *targetVehicle);
bool Player_IsInAnyVehicle(void *player);
bool Player_IsInRangeOfPoint(void *player, float range, float x, float y,
                             float z);
bool Player_PlayCrimeReport(void *player, void *suspect, int crime);
bool Player_RemoveAttachedObject(void *player, int index);
bool Player_SetAttachedObject(void *player, int index, int modelid, int bone,
                              float offsetX, float offsetY, float offsetZ,
                              float rotationX, float rotationY, float rotationZ,
                              float scaleX, float scaleY, float scaleZ,
                              uint32_t materialcolor1, uint32_t materialcolor2);
bool Player_GetAttachedObject(void *player, int index, int *modelid, int *bone,
                              float *offsetX, float *offsetY, float *offsetZ,
                              float *rotationX, float *rotationY,
                              float *rotationZ, float *scaleX, float *scaleY,
                              float *scaleZ, int *materialcolor1,
                              int *materialcolor2);
bool Player_SetFacingAngle(void *player, float angle);
bool Player_SetMarkerForPlayer(void *player, void *other, uint32_t color);
uint32_t Player_GetMarkerForPlayer(void *player, void *other);
bool Player_AllowTeleport(void *player, bool allow);
bool Player_IsTeleportAllowed(void *player);
bool Player_DisableRemoteVehicleCollisions(void *player, bool disable);
float Player_GetCameraZoom(void *player);
bool Player_SelectTextDraw(void *player, uint32_t hoverColour);
bool Player_CancelSelectTextDraw(void *player);
bool Player_SendClientCheck(void *player, int actionType, int address,
                            int offset, int count);
bool Player_Spawn(void *player);
bool Player_GPCI(void *player, struct CAPIStringView *gpci);
bool Player_IsAdmin(void *player);
bool Player_Kick(void *player);
bool Player_ShowGameText(void *player, const char *text, int time, int style);
bool Player_HideGameText(void *player, int style);
bool Player_HasGameText(void *player, int style);
bool Player_GetGameText(void *player, int style, struct CAPIStringView *message,
                        int *time, int *remaining);
bool Player_Ban(void *player);
bool Player_BanEx(void *player, const char *reason);
bool Player_SendDeathMessage(void *player, void *killer, void *killee,
                             int weapon);
bool Player_SendMessageToPlayer(void *player, void *sender,
                                const char *message);
int Player_GetVersion(void *player, struct CAPIStringView *version);
int Player_GetSkillLevel(void *player, int skill);
float Player_GetZAim(void *player);
bool Player_GetSurfingOffsets(void *player, float *offsetX, float *offsetY,
                              float *offsetZ);
bool Player_GetRotationQuat(void *player, float *x, float *y, float *z,
                            float *w);
int Player_GetPlayerSpectateID(void *player);
int Player_GetSpectateType(void *player);
uint32_t Player_GetRawIp(void *player);
bool Player_SetGravity(void *player, float gravity);
float Player_GetGravity(void *player);
bool Player_SetAdmin(void *player, bool set);
bool Player_IsSpawned(void *player);
bool Player_IsControllable(void *player);
bool Player_IsCameraTargetEnabled(void *player);
bool Player_ToggleGhostMode(void *player, bool toggle);
bool Player_GetGhostMode(void *player);
bool Player_AllowWeapons(void *player, bool allow);
bool Player_AreWeaponsAllowed(void *player);
bool Player_IsPlayerUsingOfficialClient(void *player);
int Player_GetAnimationFlags(void *player);
bool Player_IsInDriveByMode(void *player);
bool Player_IsCuffed(void *player);
bool Player_IsUsingOmp(void *player);
bool Player_IsInModShop(void *player);
int Player_GetSirenState(void *player);
int Player_GetLandingGearState(void *player);
uint32_t Player_GetHydraReactorAngle(void *player);
float Player_GetTrainSpeed(void *player);

#endif
