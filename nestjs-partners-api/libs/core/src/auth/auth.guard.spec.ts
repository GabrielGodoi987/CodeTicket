import { ConfigService } from '@nestjs/config';
import { AuthGuard } from './auth.guard';

describe('AuthGuard', () => {
  let configService: ConfigService;

  beforeAll(() => {
    configService = new ConfigService();
  });

  it('should be defined', () => {
    expect(new AuthGuard(configService)).toBeDefined();
  });
});
