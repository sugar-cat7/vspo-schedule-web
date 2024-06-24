import { v4 as uuidv4 } from 'uuid'
import { MiddlewareHandler } from "hono";
import { AppContext, HonoEnv } from "@/pkg/hono";
import { env } from 'hono/adapter'
import { type Env, zEnv } from '@/pkg/env'
import { AppLogger } from "@/pkg/logging";
import { trace } from '@opentelemetry/api';

export function init(): MiddlewareHandler<HonoEnv> {
    return async (c, next) => {
        const honoEnv = env<Env, AppContext>(c)
        const envResult = zEnv.safeParse(honoEnv)
        if (!envResult.success) {
            console.error('Failed to parse environment variables', envResult.error)
            process.exit(1)
            return
        }
        const requestId = uuidv4();
        c.set("requestId", requestId);
        const logger = new AppLogger({
            env: envResult.data,
            requestId: requestId,
        });
        c.set("services", {
            logger: logger,
            tracer: trace.getTracer('OTelCFWorkers:Fetcher')
        });
        c.set("requestUrl", envResult.data.API_BASE_URL + c.req.path);

        logger.info("[Request started]");
        await next();
    };
}