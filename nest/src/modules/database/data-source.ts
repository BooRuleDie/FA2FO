import 'dotenv/config';
import { DataSource, DataSourceOptions } from 'typeorm';

export const dataSourceOptions: DataSourceOptions = {
  // ... (keep all your existing options here)
  type: 'mysql',
  host: process.env.DB_HOST,
  port: parseInt(process.env.DB_PORT || '3306', 10),
  username: process.env.DB_USERNAME,
  password: process.env.DB_PASSWORD,
  database: process.env.DB_NAME,
  entities: [__dirname + '/../**/*.entity{.ts,.js}'],
  migrations: [__dirname + '/../../migrations/*{.ts,.js}'],
  synchronize: false,
  logging: true,
};

const dataSource = new DataSource(dataSourceOptions);

export async function runMigrations() {
  console.log('Initializing data source...');
  await dataSource.initialize();
  if (dataSource.isInitialized) {
    console.log('Data source initialized, running migrations...');
    await dataSource.runMigrations();
    console.log('Migrations are finished.');
    await dataSource.destroy();
    console.log('Data source destroyed.');
  } else {
    console.error('Data source initialization failed.');
  }
}

export default dataSource;
