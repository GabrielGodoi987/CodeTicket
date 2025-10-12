/*
  Warnings:

  - A unique constraint covering the columns `[ticketId]` on the table `Spots` will be added. If there are existing duplicate values, this will fail.
  - Added the required column `ticketId` to the `Spots` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE `Spots` ADD COLUMN `ticketId` VARCHAR(191) NOT NULL;

-- CreateTable
CREATE TABLE `ReservationHistory` (
    `id` VARCHAR(191) NOT NULL,
    `email` VARCHAR(191) NOT NULL,
    `spotId` VARCHAR(191) NOT NULL,
    `ticketKind` ENUM('full', 'half') NOT NULL,
    `createdAt` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updatedAt` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),

    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- CreateTable
CREATE TABLE `Tickets` (
    `id` VARCHAR(191) NOT NULL,
    `email` VARCHAR(191) NOT NULL,
    `spotsId` VARCHAR(191) NOT NULL,
    `ticketKind` ENUM('full', 'half') NOT NULL,
    `createdAt` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updatedAt` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),

    UNIQUE INDEX `Tickets_spotsId_key`(`spotsId`),
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- CreateIndex
CREATE UNIQUE INDEX `Spots_ticketId_key` ON `Spots`(`ticketId`);

-- AddForeignKey
ALTER TABLE `ReservationHistory` ADD CONSTRAINT `ReservationHistory_spotId_fkey` FOREIGN KEY (`spotId`) REFERENCES `Spots`(`id`) ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE `Tickets` ADD CONSTRAINT `Tickets_spotsId_fkey` FOREIGN KEY (`spotsId`) REFERENCES `Spots`(`id`) ON DELETE RESTRICT ON UPDATE CASCADE;
