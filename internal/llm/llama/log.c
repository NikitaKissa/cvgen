#include "llama.h"

static void cvgen_suppress_llama_log(enum ggml_log_level level, const char * text, void * user_data) {
	(void) level;
	(void) text;
	(void) user_data;
}


__attribute__((constructor)) static void cvgen_disable_llama_log(void) {
	llama_log_set(cvgen_suppress_llama_log, NULL);
}