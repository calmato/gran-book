import { CorsOptions } from 'cors'

const corsOptions: CorsOptions = {
  origin: ['*'],
  methods: ['GET', 'POST', 'PUT', 'PATCH', 'DELETE', 'OPTIONS'],
  allowedHeaders: ['*'],
  maxAge: 1440, // 60m * 24h
}

export { corsOptions }
