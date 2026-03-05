#ifndef TEXTDRAW_H
#define TEXTDRAW_H

#include "main.h"
#include "ompcapi.h"
#include <stdbool.h>
#include <stdint.h>

void *TextDraw_Create(float x, float y, const char *text, int *id);
bool TextDraw_Destroy(void *textdraw);
void *TextDraw_FromID(int textdrawid);
int TextDraw_GetID(void *textdraw);
bool TextDraw_IsValid(void *textdraw);
bool TextDraw_IsVisibleForPlayer(void *player, void *textdraw);
bool TextDraw_SetLetterSize(void *textdraw, float sizeX, float sizeY);
bool TextDraw_SetTextSize(void *textdraw, float sizeX, float sizeY);
bool TextDraw_SetAlignment(void *textdraw, int alignment);
bool TextDraw_SetColor(void *textdraw, uint32_t color);
bool TextDraw_SetUseBox(void *textdraw, bool use);
bool TextDraw_SetBoxColor(void *textdraw, uint32_t color);
bool TextDraw_SetShadow(void *textdraw, int size);
bool TextDraw_SetOutline(void *textdraw, int size);
bool TextDraw_SetBackgroundColor(void *textdraw, uint32_t color);
bool TextDraw_SetFont(void *textdraw, int font);
bool TextDraw_SetProportional(void *textdraw, bool set);
bool TextDraw_SetSelectable(void *textdraw, bool set);
bool TextDraw_ShowForPlayer(void *player, void *textdraw);
bool TextDraw_HideForPlayer(void *player, void *textdraw);
bool TextDraw_ShowForAll(void *textdraw);
bool TextDraw_HideForAll(void *textdraw);
bool TextDraw_SetString(void *textdraw, const char *text);
bool TextDraw_SetPreviewModel(void *textdraw, int model);
bool TextDraw_SetPreviewRot(void *textdraw, float rotationX, float rotationY,
                            float rotationZ, float zoom);
bool TextDraw_SetPreviewVehCol(void *textdraw, int color1, int color2);
bool TextDraw_SetPos(void *textdraw, float x, float y);
bool TextDraw_GetString(void *textdraw, struct CAPIStringView *text);
bool TextDraw_GetLetterSize(void *textdraw, float *sizeX, float *sizeY);
bool TextDraw_GetTextSize(void *textdraw, float *sizeX, float *sizeY);
bool TextDraw_GetPos(void *textdraw, float *x, float *y);
int TextDraw_GetColor(void *textdraw);
int TextDraw_GetBoxColor(void *textdraw);
int TextDraw_GetBackgroundColor(void *textdraw);
int TextDraw_GetShadow(void *textdraw);
int TextDraw_GetOutline(void *textdraw);
int TextDraw_GetFont(void *textdraw);
bool TextDraw_IsBox(void *textdraw);
bool TextDraw_IsProportional(void *textdraw);
bool TextDraw_IsSelectable(void *textdraw);
int TextDraw_GetAlignment(void *textdraw);
int TextDraw_GetPreviewModel(void *textdraw);
bool TextDraw_GetPreviewRot(void *textdraw, float *x, float *y, float *z,
                            float *zoom);
bool TextDraw_GetPreviewVehColor(void *textdraw, int *color1, int *color2);
bool TextDraw_SetStringForPlayer(void *textdraw, void *player,
                                 const char *text);

#endif
