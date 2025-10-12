/*
  Warnings:

  - You are about to drop the column `ticketId` on the `Spots` table. All the data in the column will be lost.
  - A unique constraint covering the columns `[name]` on the table `Spots` will be added. If there are existing duplicate values, this will fail.

*/
-- DropIndex
DROP INDEX `Spots_ticketId_name_key` ON `Spots`;

-- AlterTable
ALTER TABLE `Spots` DROP COLUMN `ticketId`;

-- CreateIndex
CREATE UNIQUE INDEX `Spots_name_key` ON `Spots`(`name`);
