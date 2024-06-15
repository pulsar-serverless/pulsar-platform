export interface Project {
  id: string;
  name: string;
  subdomain: string;
  deploymentStatus: 'none' | 'failed' | 'done' | 'building';
  createdAt: string;
  updatedAt: string;
}
