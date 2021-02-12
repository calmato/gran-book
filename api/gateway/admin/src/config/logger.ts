import { Configuration } from 'log4js'

const logDir: string = process.env.LOG_PATH || '/var/log/gateway'

const loggerConfig: Configuration = {
  appenders: {
    consoleLog: {
      type: 'console'
    },
    systemLog: {
      type: 'file',
      filename: `${logDir}/system.log`,
      // maxLogSize: 5000000, // 5MB
      // backups: 5, // 世代管理は5ファイルまで、古いやつgzで圧縮されていく
      // compress: true
    },
    accessLog: {
      type: "dateFile",
      filename: `${logDir}/access/access.log`,
      // pattern: "yyyy-MM-dd", // 日毎にファイル分割
      // daysToKeep: 5, // 5日分の世代管理設定
      // compress: true,
      // keepFileExt: true,
    }
  },
  categories: {
    default: {
      appenders: ['consoleLog'],
      level: 'ALL'
    },
    system: {
      appenders: ['systemLog'],
      level: 'ERROR'
    },
    access: {
      appenders: ["accessLog"],
      level: "DEBUG"
    }
  }
}

export { loggerConfig }
