import { NestFactory } from '@nestjs/core';
import { SeederModule } from './seeder.module';
import { SeederService } from './seeder.service';
import { Logger } from '@nestjs/common';

async function bootstrap() {
  const logger = new Logger('Seeder');
  // We create a standalone application context, not a web server
  const appContext = await NestFactory.createApplicationContext(SeederModule);

  logger.log('Seeder application context created.');

  const seeder = appContext.get(SeederService);

  try {
    await seeder.seed();
    logger.log('Seeding complete!');
  } catch (error) {
    logger.error('Seeding failed!');
    logger.error(error);
  } finally {
    await appContext.close();
  }
}

bootstrap().catch((err) => {
  console.error('Seeding failed:', err);
  process.exit(1);
});
