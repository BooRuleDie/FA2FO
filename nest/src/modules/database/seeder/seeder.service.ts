import { Injectable, Logger } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { User, UserRoles } from '../entities/user.entity';
import { Post } from '../entities/post.entity';
import { Like } from '../entities/like.entity';
import { faker } from '@faker-js/faker';
import * as bcrypt from 'bcrypt';

@Injectable()
export class SeederService {
  private readonly logger = new Logger(SeederService.name);
  private readonly saltRounds = 10;

  constructor(
    @InjectRepository(User)
    private readonly userRepository: Repository<User>,
    @InjectRepository(Post)
    private readonly postRepository: Repository<Post>,
    @InjectRepository(Like)
    private readonly likeRepository: Repository<Like>,
  ) {}

  async seed() {
    this.logger.log('Starting the seeding process...');

    await this.clearDatabase();
    const users = await this.seedUsers(10);
    const posts = await this.seedPosts(users, 50);
    await this.seedLikes(users, posts);

    this.logger.log('Seeding completed successfully.');
  }

  private async clearDatabase() {
    this.logger.log('Clearing database...');
    await this.likeRepository.createQueryBuilder().delete().execute();
    await this.postRepository.createQueryBuilder().delete().execute();
    await this.userRepository.createQueryBuilder().delete().execute();
    this.logger.log('Database cleared.');
  }

  private async seedUsers(count: number): Promise<User[]> {
    const users: User[] = [];
    const hashedPassword = await bcrypt.hash(
      process.env.SEED_USER_PASSWD || 'defaultpassword',
      this.saltRounds,
    );

    for (let i = 0; i < count; i++) {
      // insert the admin user on first INSERT
      if (i === 0) {
        const user = this.userRepository.create({
          email: process.env.ADMIN_EMAIL,
          password: hashedPassword,
          username: process.env.ADMIN_USERNAME,
          role: UserRoles.ADMIN,
        });
        users.push(user);
        continue;
      }

      const user = this.userRepository.create({
        email: faker.internet.email(),
        password: hashedPassword,
        username: faker.person.firstName(),
        role: UserRoles.CUSTOMER,
      });
      users.push(user);
    }
    this.logger.log(`Seeding ${users.length} users...`);
    return this.userRepository.save(users);
  }

  private async seedPosts(users: User[], count: number): Promise<Post[]> {
    const posts: Post[] = [];
    for (let i = 0; i < count; i++) {
      const post = this.postRepository.create({
        title: faker.lorem.sentence(),
        content: faker.lorem.paragraphs(3),
        author: users[Math.floor(Math.random() * users.length)], // Assign a random author
      });
      posts.push(post);
    }
    this.logger.log(`Seeding ${posts.length} posts...`);
    return this.postRepository.save(posts);
  }

  private async seedLikes(users: User[], posts: Post[]) {
    const likes: Like[] = [];
    for (const user of users) {
      // Each user likes a random number of posts
      const postsToLike = faker.helpers.arrayElements(
        posts,
        faker.number.int({ min: 5, max: 20 }),
      );
      for (const post of postsToLike) {
        // Prevent user from liking their own post for more realistic data
        if (post.author.id !== user.id) {
          const like = this.likeRepository.create({
            user: user,
            post: post,
          });
          likes.push(like);
        }
      }
    }
    this.logger.log(`Seeding ${likes.length} likes...`);
    // Use chunking for better performance on bulk inserts
    await this.likeRepository.save(likes, { chunk: 100 });
  }
}
