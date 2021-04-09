import multer from 'multer'

const upload: multer.Multer = multer({ dest: 'tmp/uploads' })

export default upload
