import multer from 'multer'

const uploadPath: string = process.env.TEMPORARY_UPLOAD_PATH || 'tmp/uploads'
const upload: multer.Multer = multer({ dest: uploadPath })

export default upload
