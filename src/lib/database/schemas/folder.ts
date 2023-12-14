import { pgTable, text, timestamp, uuid } from "drizzle-orm/pg-core"
import WorkSpace from "./workspace"

const Folders = pgTable("folders", {
  id: uuid("id").defaultRandom().primaryKey(),
  createdAt: timestamp("created_at", { withTimezone: true, mode: "string" }),
  title: text("title").notNull(),
  iconId: text("icon_id").notNull(),
  data: text("data"),
  inTrash: text("in_trash"),
  logo: text("logo"),
  bannerUrl: text("banner_url").notNull(),
  workspaceId: uuid("workspace_id").references(() => WorkSpace.id, {
    onDelete: "cascade",
  }),
})

export default Folders
