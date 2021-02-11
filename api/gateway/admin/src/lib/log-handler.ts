/* eslint-disable @typescript-eslint/ban-types */
/* eslint-disable @typescript-eslint/no-unsafe-assignment */
/* eslint-disable @typescript-eslint/unbound-method */
/* eslint-disable prefer-rest-params */
import { Request, Response, NextFunction} from 'express'
import logger from '~/plugins/logger'

export function logHandler(req: Request, res: Response, next: NextFunction): void {
  const [oldWrite, oldEnd] = [res.write, res.end];
  const chunks: Buffer[] = [];

  (res.write as unknown) = (chunk: any) => {
    chunks.push(Buffer.from(chunk));
    (oldWrite as Function).apply(res, arguments);
  };

  res.end = (chunk: any) => {
    if (chunk) {
      chunks.push(Buffer.from(chunk));
    }

    const body = Buffer.concat(chunks).toString('utf8');

    const logs: any = {
      requestHeader: req.headers,
      requestBody: req.body,
      responseBody: JSON.parse(body),
    };

    logger.access.info(req.method, req.path, logs);

    (oldEnd as Function).apply(res, arguments);
  };

  next()
}
/* eslint-enable */
