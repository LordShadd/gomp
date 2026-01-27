#ifndef EVENTS_H
#define EVENTS_H

#include "ompcapi.h"
#include <stdbool.h>
#include <stdint.h>

bool Event_AddHandler(const char *name, int priority, void *callback);
bool Event_RemoveHandler(const char *name, int priority, void *callback);
bool Event_RemoveAllHandlers(const char *name, int priority);

#endif
