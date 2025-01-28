import { env } from "hono/adapter";
import { createMiddleware } from "hono/factory";
import type { AppEnv } from "../../../../config/env";
import { AppLogger } from "../../../../pkg/logging";
import type { AppContext } from "../app";
import type { HonoEnv } from "../env";

export const init = createMiddleware<HonoEnv>(async (c, next) => {
  const honoEnv = env<AppEnv, AppContext>(c);
  const requestId = c.get("requestId");
  const logger = new AppLogger({
    env: honoEnv,
    requestId: requestId,
  });
  c.set("requestId", requestId);
  c.set("services", {
    logger: logger,
  });
  logger.info("[Request started]");
  await next();
});
