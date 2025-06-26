import { MigrationInterface, QueryRunner } from 'typeorm';

export class UserTableUpdate1750938194631 implements MigrationInterface {
  name = 'UserTableUpdate1750938194631';

  public async up(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.query(
      `ALTER TABLE \`users\` ADD \`password\` varchar(255) NOT NULL`,
    );
    await queryRunner.query(
      `ALTER TABLE \`users\` ADD \`role\` enum ('admin', 'customer') NOT NULL DEFAULT 'customer'`,
    );
  }

  public async down(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.query(`ALTER TABLE \`users\` DROP COLUMN \`role\``);
    await queryRunner.query(`ALTER TABLE \`users\` DROP COLUMN \`password\``);
  }
}
