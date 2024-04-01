export interface Project {
  id: string;
  name: string;
  deploymentStatus: 'none' | 'failed' | 'done' | 'building';
  createdAt: string;
  updatedAt: string;
}
