import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { SeederService } from './seeder.service';
import { User } from '../entities/user.entity';
import { Post } from '../entities/post.entity';
import { Like } from '../entities/like.entity';
import { DatabaseModule } from '../database.module';

@Module({
  imports: [
    // We need the DatabaseModule to get the TypeORM configuration
    DatabaseModule,
    // Import all the entities we need to seed
    TypeOrmModule.forFeature([User, Post, Like]),
  ],
  providers: [SeederService],
})
export class SeederModule {}
