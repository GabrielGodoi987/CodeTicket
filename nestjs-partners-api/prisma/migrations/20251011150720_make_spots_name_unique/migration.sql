/*
  Warnings:

  - A unique constraint covering the columns `[ticketId,name]` on the table `Spots` will be added. If there are existing duplicate values, this will fail.

*/
-- DropIndex
DROP INDEX `Spots_ticketId_key` ON `Spots`;

-- CreateIndex
CREATE UNIQUE INDEX `Spots_ticketId_name_key` ON `Spots`(`ticketId`, `name`);
