import { Module } from '@nestjs/common';
import { EventsService } from './events.service';

@Module({
  controllers: [],
  providers: [EventsService],
  exports: [EventsService],
})
export class EventsCoreModule {}
