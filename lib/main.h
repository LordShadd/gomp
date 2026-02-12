#ifndef MAIN_H
#define MAIN_H

#include "ompcapi.h"
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

extern struct OMPAPI_t api;

void setComponentVersion(uint8_t major, uint8_t minor, uint8_t patch,
                         uint16_t prerel);
void setComponentName(char *name);

extern void onReady();
extern void onReset();
extern void onFree();

extern bool onPlayerConnect(struct EventArgs_onPlayerConnect *args);
extern bool onPlayerDisconnect(struct EventArgs_onPlayerDisconnect *args);
extern bool onPlayerRequestClass(struct EventArgs_onPlayerRequestClass *args);
extern bool onPlayerRequestSpawn(struct EventArgs_onPlayerRequestSpawn *args);
extern bool onPlayerSpawn(struct EventArgs_onPlayerSpawn *args);
extern bool onPlayerDeath(struct EventArgs_onPlayerDeath *args);
extern bool onPlayerUpdate(struct EventArgs_onPlayerUpdate *args);
extern bool onPlayerText(struct EventArgs_onPlayerText *args);
extern bool onPlayerCommandText(struct EventArgs_onPlayerCommandText *args);
extern bool
onPlayerInteriorChange(struct EventArgs_onPlayerInteriorChange *args);
extern bool onPlayerStateChange(struct EventArgs_onPlayerStateChange *args);
extern bool
onPlayerKeyStateChange(struct EventArgs_onPlayerKeyStateChange *args);
extern bool onPlayerEnterVehicle(struct EventArgs_onPlayerEnterVehicle *args);
extern bool onPlayerExitVehicle(struct EventArgs_onPlayerExitVehicle *args);
extern bool
onPlayerEnterCheckpoint(struct EventArgs_onPlayerEnterCheckpoint *args);
extern bool
onPlayerLeaveCheckpoint(struct EventArgs_onPlayerLeaveCheckpoint *args);
extern bool
onPlayerEnterRaceCheckpoint(struct EventArgs_onPlayerEnterRaceCheckpoint *args);
extern bool
onPlayerLeaveRaceCheckpoint(struct EventArgs_onPlayerLeaveRaceCheckpoint *args);
extern bool onPlayerGiveDamage(struct EventArgs_onPlayerGiveDamage *args);
extern bool onPlayerTakeDamage(struct EventArgs_onPlayerTakeDamage *args);
extern bool
onPlayerGiveDamageActor(struct EventArgs_onPlayerGiveDamageActor *args);
extern bool onPlayerShotMissed(struct EventArgs_onPlayerShotMissed *args);
extern bool onPlayerShotObject(struct EventArgs_onPlayerShotObject *args);
extern bool onPlayerShotPlayer(struct EventArgs_onPlayerShotPlayer *args);
extern bool
onPlayerShotPlayerObject(struct EventArgs_onPlayerShotPlayerObject *args);
extern bool onPlayerShotVehicle(struct EventArgs_onPlayerShotVehicle *args);
extern bool onPlayerPickUpPickup(struct EventArgs_onPlayerPickUpPickup *args);
extern bool onPlayerObjectMove(struct EventArgs_onPlayerObjectMove *args);
extern bool onPlayerEditObject(struct EventArgs_onPlayerEditObject *args);
extern bool
onPlayerEditAttachedObject(struct EventArgs_onPlayerEditAttachedObject *args);
extern bool onPlayerSelectObject(struct EventArgs_onPlayerSelectObject *args);
extern bool onPlayerClickMap(struct EventArgs_onPlayerClickMap *args);
extern bool onPlayerClickTextDraw(struct EventArgs_onPlayerClickTextDraw *args);
extern bool
onPlayerClickPlayerTextDraw(struct EventArgs_onPlayerClickPlayerTextDraw *args);
extern bool onPlayerClickPlayer(struct EventArgs_onPlayerClickPlayer *args);
extern bool onPlayerStreamIn(struct EventArgs_onPlayerStreamIn *args);
extern bool onPlayerStreamOut(struct EventArgs_onPlayerStreamOut *args);
extern bool onPlayerExitedMenu(struct EventArgs_onPlayerExitedMenu *args);
extern bool
onPlayerSelectedMenuRow(struct EventArgs_onPlayerSelectedMenuRow *args);
extern bool
onPlayerRequestDownload(struct EventArgs_onPlayerRequestDownload *args);

extern void onTick();
extern bool onIncomingConnection(struct EventArgs_onIncomingConnection *args);
extern bool onRconLoginAttempt(struct EventArgs_onRconLoginAttempt *args);
extern bool onConsoleText(struct EventArgs_onConsoleText *args);
extern bool onDialogResponse(struct EventArgs_onDialogResponse *args);

extern bool onVehicleSpawn(struct EventArgs_onVehicleSpawn *args);
extern bool onVehicleDeath(struct EventArgs_onVehicleDeath *args);
extern bool onVehicleMod(struct EventArgs_onVehicleMod *args);
extern bool onVehiclePaintJob(struct EventArgs_onVehiclePaintJob *args);
extern bool onVehicleRespray(struct EventArgs_onVehicleRespray *args);
extern bool
onVehicleDamageStatusUpdate(struct EventArgs_onVehicleDamageStatusUpdate *args);
extern bool
onVehicleSirenStateChange(struct EventArgs_onVehicleSirenStateChange *args);
extern bool onVehicleStreamIn(struct EventArgs_onVehicleStreamIn *args);
extern bool onVehicleStreamOut(struct EventArgs_onVehicleStreamOut *args);
extern bool
onUnoccupiedVehicleUpdate(struct EventArgs_onUnoccupiedVehicleUpdate *args);
extern bool onTrailerUpdate(struct EventArgs_onTrailerUpdate *args);

extern bool onActorStreamIn(struct EventArgs_onActorStreamIn *args);
extern bool onActorStreamOut(struct EventArgs_onActorStreamOut *args);
extern bool onObjectMove(struct EventArgs_onObjectMove *args);
extern bool onEnterExitModShop(struct EventArgs_onEnterExitModShop *args);

extern void entryPoint();

#endif
