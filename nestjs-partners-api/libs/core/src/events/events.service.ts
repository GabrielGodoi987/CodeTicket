import { Injectable } from '@nestjs/common';
import { CreateEventDto } from './dto/create-event.dto';
import { UpdateEventDto } from './dto/update-event.dto';
import { PrismaService } from '@app/core/prisma/prisma.service';
import { ReserveSpotDto } from './dto/reserve-spot.dto';
import { Prisma, SpotStatus } from '@prisma/client';

@Injectable()
export class EventsService {
  constructor(private prismaService: PrismaService) {}
  create(createEventDto: CreateEventDto) {
    return this.prismaService.events.create({
      data: createEventDto,
    });
  }

  findAll() {
    return this.prismaService.events.findMany();
  }

  findOne(id: string) {
    return this.prismaService.events.findUnique({
      where: {
        id,
      },
    });
  }

  update(id: string, updateEventDto: UpdateEventDto) {
    return this.prismaService.events.update({
      where: {
        id,
      },
      data: updateEventDto,
    });
  }

  remove(id: string) {
    return this.prismaService.events.delete({
      where: {
        id,
      },
    });
  }

  async reserveSpot(reserveSpotDto: ReserveSpotDto) {
    const { spots, eventId, ticket_kind, email } = reserveSpotDto;

    const foundedSpots = await this.prismaService.spots.findMany({
      where: {
        eventId,
        name: {
          in: spots,
        },
      },
    });

    if (foundedSpots.length !== spots.length) {
      const foundSpotsName = foundedSpots.map((spot) => spot.name);

      const notFoundSpotsName = spots.filter(
        (spotName) => !foundSpotsName.includes(spotName),
      );

      throw new Error(`Spots ${notFoundSpotsName.join(', ')} not found`);
    }

    try {
      const tickets = await this.prismaService.$transaction(
        async (prisma) => {
          await prisma.reservationHistory.createMany({
            data: foundedSpots.map((spots) => ({
              spotId: spots.id,
              ticketKind: ticket_kind,
              email,
            })),
          });

          await prisma.spots.updateMany({
            where: {
              id: {
                in: foundedSpots.map((spot) => spot.id),
              },
            },
            data: {
              status: SpotStatus.RESERVED,
            },
          });

          const tickets = await Promise.all(
            foundedSpots.map((spot) =>
              prisma.tickets.create({
                data: {
                  spotsId: spot.id,
                  ticketKind: ticket_kind,
                  email,
                },
              }),
            ),
          );

          return tickets;
        },
        { isolationLevel: Prisma.TransactionIsolationLevel.ReadCommitted },
      );
      return tickets;
    } catch (error) {
      if (error instanceof Prisma.PrismaClientKnownRequestError) {
        switch (error.code) {
          case 'P2002': // unique constraint violation
          case 'P2034': // transaction conflict
            throw new Error('Some spots are already reserved');
        }
      }

      throw error;
    }
  }
}
