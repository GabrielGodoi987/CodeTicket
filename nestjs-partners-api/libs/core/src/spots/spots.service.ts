import { Injectable, NotFoundException } from '@nestjs/common';
import { CreateSpotDto } from './dto/create-spot.dto';
import { UpdateSpotDto } from './dto/update-spot.dto';
import { PrismaService } from '@app/core/prisma/prisma.service';
import { SpotStatus } from '@prisma/client';

@Injectable()
export class SpotsService {
  constructor(private prismaService: PrismaService) {}

  async create(createSpotDto: CreateSpotDto & { eventId: string }) {
    const event = await this.prismaService.events.findFirst({
      where: {
        id: createSpotDto.eventId,
      },
    });

    if (!event) {
      throw new Error('Event not found');
    }

    return this.prismaService.spots.create({
      data: {
        ...createSpotDto,
        status: SpotStatus.AVAILABLE,
      },
    });
  }

  async findAll(eventId: string) {
    return await this.prismaService.spots.findMany({
      where: {
        eventId,
      },
    });
  }

  async findOne(eventId: string, spotId: string) {
    return await this.prismaService.spots.findFirst({
      where: {
        id: spotId,
        eventId,
      },
    });
  }

  async update(eventId: string, spotId: string, updateSpotDto: UpdateSpotDto) {
    const spot = await this.findOne(eventId, spotId);
    if (!spot)
      throw new NotFoundException('Spot not found or does not belong to event');

    return await this.prismaService.spots.update({
      where: {
        id: spotId,
      },
      data: updateSpotDto,
    });
  }

  async remove(eventId: string, spotId: string) {
    const spot = await this.findOne(eventId, spotId);
    if (!spot)
      throw new NotFoundException('Spot not found or does not belong to event');

    return await this.prismaService.spots.delete({
      where: {
        id: spotId,
      },
    });
  }
}
