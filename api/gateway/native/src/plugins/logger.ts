import log4js from 'log4js'
import { loggerConfig } from '~/config/logger'

log4js.configure(loggerConfig)

const logger = {
  default: log4js.getLogger('default'),
  system: log4js.getLogger('system'),
  access: log4js.getLogger('access'),
}

export default logger
