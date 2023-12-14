import { pgTable, text, timestamp, uuid } from "drizzle-orm/pg-core"

const WorkSpaces = pgTable("workspaces", {
  id: uuid("id").defaultRandom().primaryKey(),
  createdAt: timestamp("created_at", { withTimezone: true, mode: "string" }),
  workspaceOwner: uuid("workspace_owner").notNull(),
  title: text("title").notNull(),
  iconId: text("icon_id").notNull(),
  data: text("data"),
  inTrash: text("in_trash"),
  logo: text("logo"),
  bannerUrl: text("banner_url").notNull(),
})

export default WorkSpaces
