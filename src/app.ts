import express from "express"
import { Express } from "express-serve-static-core"
import { readFileSync } from "fs"
import MarkdownIt from "markdown-it"
const md = new MarkdownIt({
    html: true,
    linkify: true,
    typographer: true,
})

class App {
    public express: Express
    fileName: String

    constructor() {
        this.fileName = "README.md"
        this.express = express()
        this.express.set("view engine", "ejs")
        this.mountRoutes()
    }

    private readMarkdown(fileName:String): string {
        let data = readFileSync(fileName.toString(), "utf8")
        return data
    }

    private mountRoutes(): void {
        const router = express.Router()
        router.get("/", (req, res) => {
            res.render("md", {
                content: md.render(this.readMarkdown("README.md")),
            })
        })
        this.express.use("/", router)
    }
}

export default new App().express
