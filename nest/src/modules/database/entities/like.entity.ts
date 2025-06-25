import {
  Entity,
  PrimaryGeneratedColumn,
  CreateDateColumn,
  ManyToOne,
  JoinColumn,
  Unique,
} from 'typeorm';
import { User } from './user.entity';
import { Post } from './post.entity';

@Entity({ name: 'likes' })
@Unique(['user', 'post']) // Ensures a user can only like a post once.
export class Like {
  @PrimaryGeneratedColumn()
  id: number;

  // Establishes the relationship with the User entity.
  // When a User is deleted, their likes are also deleted (onDelete: 'CASCADE').
  @ManyToOne(() => User, {
    nullable: false,
    onDelete: 'CASCADE',
  })
  @JoinColumn({ name: 'user_id' })
  user: User;

  // Establishes the relationship with the Post entity.
  // When a Post is deleted, all its likes are also deleted.
  @ManyToOne(() => Post, {
    nullable: false,
    onDelete: 'CASCADE',
  })
  @JoinColumn({ name: 'post_id' })
  post: Post;

  @CreateDateColumn({ name: 'created_at', type: 'timestamp' })
  createdAt: Date;
}
