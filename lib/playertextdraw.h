#ifndef PLAYERTEXTDRAW_H
#define PLAYERTEXTDRAW_H

#include "main.h"
#include "ompcapi.h"
#include <stdbool.h>
#include <stdint.h>

void *PlayerTextDraw_Create(void *player, float x, float y, const char *text,
                             int *id);
bool PlayerTextDraw_Destroy(void *player, void *textdraw);
void *PlayerTextDraw_FromID(void *player, int textdrawid);
int PlayerTextDraw_GetID(void *player, void *textdraw);
bool PlayerTextDraw_IsValid(void *player, void *textdraw);
bool PlayerTextDraw_IsVisible(void *player, void *textdraw);
bool PlayerTextDraw_SetLetterSize(void *player, void *textdraw, float x,
                                  float y);
bool PlayerTextDraw_SetTextSize(void *player, void *textdraw, float x, float y);
bool PlayerTextDraw_SetAlignment(void *player, void *textdraw, int alignment);
bool PlayerTextDraw_SetColor(void *player, void *textdraw, uint32_t color);
bool PlayerTextDraw_UseBox(void *player, void *textdraw, bool use);
bool PlayerTextDraw_SetBoxColor(void *player, void *textdraw, uint32_t color);
bool PlayerTextDraw_SetShadow(void *player, void *textdraw, int size);
bool PlayerTextDraw_SetOutline(void *player, void *textdraw, int size);
bool PlayerTextDraw_SetBackgroundColor(void *player, void *textdraw,
                                       uint32_t color);
bool PlayerTextDraw_SetFont(void *player, void *textdraw, int font);
bool PlayerTextDraw_SetProportional(void *player, void *textdraw, bool set);
bool PlayerTextDraw_SetSelectable(void *player, void *textdraw, bool set);
bool PlayerTextDraw_Show(void *player, void *textdraw);
bool PlayerTextDraw_Hide(void *player, void *textdraw);
bool PlayerTextDraw_SetString(void *player, void *textdraw, const char *text);
bool PlayerTextDraw_SetPreviewModel(void *player, void *textdraw, int model);
bool PlayerTextDraw_SetPreviewRot(void *player, void *textdraw, float rx,
                                  float ry, float rz, float zoom);
bool PlayerTextDraw_SetPreviewVehCol(void *player, void *textdraw, int color1,
                                     int color2);
bool PlayerTextDraw_SetPos(void *player, void *textdraw, float x, float y);
bool PlayerTextDraw_GetString(void *player, void *textdraw,
                              struct CAPIStringView *text);
bool PlayerTextDraw_GetLetterSize(void *player, void *textdraw, float *x,
                                  float *y);
bool PlayerTextDraw_GetTextSize(void *player, void *textdraw, float *x,
                                float *y);
bool PlayerTextDraw_GetPos(void *player, void *textdraw, float *x, float *y);
int PlayerTextDraw_GetColor(void *player, void *textdraw);
int PlayerTextDraw_GetBoxColor(void *player, void *textdraw);
int PlayerTextDraw_GetBackgroundColor(void *player, void *textdraw);
int PlayerTextDraw_GetShadow(void *player, void *textdraw);
int PlayerTextDraw_GetOutline(void *player, void *textdraw);
int PlayerTextDraw_GetFont(void *player, void *textdraw);
bool PlayerTextDraw_IsBox(void *player, void *textdraw);
bool PlayerTextDraw_IsProportional(void *player, void *textdraw);
bool PlayerTextDraw_IsSelectable(void *player, void *textdraw);
int PlayerTextDraw_GetAlignment(void *player, void *textdraw);
int PlayerTextDraw_GetPreviewModel(void *player, void *textdraw);
bool PlayerTextDraw_GetPreviewRot(void *player, void *textdraw, float *rx,
                                  float *ry, float *rz, float *zoom);
bool PlayerTextDraw_GetPreviewVehColor(void *player, void *textdraw,
                                       int *color1, int *color2);

#endif
