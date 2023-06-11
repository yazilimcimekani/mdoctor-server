import { Express } from 'express-serve-static-core'
import server from './server'

class MDoctorServer {
  private app: Express

  constructor(app: Express) {
    this.app = app
  }

  serve(port: number): void {
    this.app.listen(port, () => {
      console.log(`Listening on port ${port}`)
    })
  }
}

export default new MDoctorServer(server)
