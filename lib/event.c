#include "event.h"
#include "main.h"

bool Event_AddHandler(const char *name, int priority, void *callback) {
  return api.Event.AddHandler(name, priority, callback);
}

bool Event_RemoveHandler(const char *name, int priority, void *callback) {
  return api.Event.RemoveHandler(name, priority, callback);
}

bool Event_RemoveAllHandlers(const char *name, int priority) {
  return api.Event.RemoveAllHandlers(name, priority);
}
