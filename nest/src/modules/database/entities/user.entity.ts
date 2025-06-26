import {
  Entity,
  PrimaryGeneratedColumn,
  Column,
  CreateDateColumn,
  UpdateDateColumn,
  OneToMany,
} from 'typeorm';
import { Post } from './post.entity';
import { Like } from './like.entity';

export enum UserRoles {
  ADMIN = 'admin',
  CUSTOMER = 'customer',
}

@Entity({ name: 'users' })
export class User {
  @PrimaryGeneratedColumn()
  id: number;

  @Column({ length: 255, unique: true })
  username: string;

  @Column({ length: 255, unique: true })
  email: string;

  @Column({ length: 255 })
  password: string;

  @Column({
    type: 'enum',
    enum: UserRoles,
    default: UserRoles.CUSTOMER,
  })
  role: UserRoles;

  // One-to-many relationship with posts
  @OneToMany(() => Post, (post) => post.author)
  posts: Post[];

  // One-to-many relationship with likes
  @OneToMany(() => Like, (like) => like.user)
  likes: Like[];

  @CreateDateColumn({ name: 'created_at', type: 'timestamp' })
  createdAt: Date;

  @UpdateDateColumn({ name: 'updated_at', type: 'timestamp' })
  updatedAt: Date;
}
