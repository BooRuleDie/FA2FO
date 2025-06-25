import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { runMigrations } from './modules/database/data-source';

async function bootstrap() {
  console.log('Starting migration process...');
  await runMigrations();
  console.log('Migration process finished. Exiting.');

  console.log('Starting NestJS application...');
  const app = await NestFactory.create(AppModule);
  // Add any global pipes, interceptors, etc. here
  await app.listen(3000);
  console.log(`Application is running on: ${await app.getUrl()}`);
}

bootstrap().catch((err) => {
  console.error('Application failed to start:', err);
  process.exit(1);
});
