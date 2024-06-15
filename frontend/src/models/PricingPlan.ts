export interface PricingPlan {
  id: string;
  name: string;
  desc: string;
  tier: string;
  allocatedMemory: number;
  allocatedBandwidth: number;
  allocatedRequests: number;
  price?: number;
}
