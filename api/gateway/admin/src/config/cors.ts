import { CorsOptions } from 'cors'

const corsOptions: CorsOptions = {
  origin: true,
  methods: ['GET', 'POST', 'PUT', 'PATCH', 'DELETE', 'OPTIONS'],
  allowedHeaders: [
    'Content-Type',
    'Accept',
    'User-Agent',
    'Authorization',
    'X-Forwarded-For',
    'X-Forwarded-Proto',
    'X-Real-Ip',
  ],
  credentials: true,
  maxAge: 1440, // 60m * 24h
  preflightContinue: true,
}

export { corsOptions }
