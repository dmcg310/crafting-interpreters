#include <stdio.h>

#include "common.h"
#include "compiler.h"
#include "scanner.h"

void compile(const char *source, Chunk *chunk) {
  initScanner(source);
  advance();
  expression();
  consume(TOKEN_EOF, "[COMPILER ERROR] Expect end of expression");
}
