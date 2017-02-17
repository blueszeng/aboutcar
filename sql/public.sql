/*
 Navicat PostgreSQL Data Transfer

 Source Server         : test
 Source Server Version : 90601
 Source Host           : localhost
 Source Database       : aboubcar
 Source Schema         : public

 Target Server Version : 90601
 File Encoding         : utf-8

 Date: 02/17/2017 13:01:18 PM
*/

-- ----------------------------
--  Table structure for edusite
-- ----------------------------
DROP TABLE IF EXISTS "public"."edusite";
CREATE TABLE "public"."edusite" (
	"uuid" varchar(255) NOT NULL COLLATE "default",
	"county" varchar(255) NOT NULL DEFAULT NULL::character varying COLLATE "default",
	"train_name" varchar(255) NOT NULL COLLATE "default"
)
WITH (OIDS=FALSE);
ALTER TABLE "public"."edusite" OWNER TO "zyg";

-- ----------------------------
--  Table structure for coach
-- ----------------------------
DROP TABLE IF EXISTS "public"."coach";
CREATE TABLE "public"."coach" (
	"uuid" varchar(255) NOT NULL COLLATE "default",
	"name" varchar(255) NOT NULL COLLATE "default",
	"phone" varchar(12) NOT NULL COLLATE "default",
	"train_id" varchar(255) NOT NULL COLLATE "default"
)
WITH (OIDS=FALSE);
ALTER TABLE "public"."coach" OWNER TO "zyg";

-- ----------------------------
--  Table structure for train
-- ----------------------------
DROP TABLE IF EXISTS "public"."train";
CREATE TABLE "public"."train" (
	"uuid" varchar(255) NOT NULL COLLATE "default",
	"project_name" varchar(255) NOT NULL DEFAULT 0 COLLATE "default",
	"price" varchar(255) NOT NULL COLLATE "default",
	"edusite_id" varchar(255) NOT NULL COLLATE "default"
)
WITH (OIDS=FALSE);
ALTER TABLE "public"."train" OWNER TO "zyg";

-- ----------------------------
--  Primary key structure for table edusite
-- ----------------------------
ALTER TABLE "public"."edusite" ADD PRIMARY KEY ("uuid") NOT DEFERRABLE INITIALLY IMMEDIATE;

-- ----------------------------
--  Indexes structure for table edusite
-- ----------------------------
CREATE UNIQUE INDEX  "UQE_edusite_uuid" ON "public"."edusite" USING btree(uuid COLLATE "default" "pg_catalog"."text_ops" ASC NULLS LAST);

-- ----------------------------
--  Primary key structure for table coach
-- ----------------------------
ALTER TABLE "public"."coach" ADD PRIMARY KEY ("uuid", "train_id") NOT DEFERRABLE INITIALLY IMMEDIATE;

-- ----------------------------
--  Indexes structure for table coach
-- ----------------------------
CREATE UNIQUE INDEX  "UQE_coach_uuid" ON "public"."coach" USING btree(uuid COLLATE "default" "pg_catalog"."text_ops" ASC NULLS LAST);

-- ----------------------------
--  Primary key structure for table train
-- ----------------------------
ALTER TABLE "public"."train" ADD PRIMARY KEY ("uuid", "edusite_id") NOT DEFERRABLE INITIALLY IMMEDIATE;

-- ----------------------------
--  Indexes structure for table train
-- ----------------------------
CREATE UNIQUE INDEX  "UQE_train_uuid" ON "public"."train" USING btree(uuid COLLATE "default" "pg_catalog"."text_ops" ASC NULLS LAST);

