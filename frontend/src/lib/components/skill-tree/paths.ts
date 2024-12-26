// Tracks the next hop you should take to traverse a shortest path from any
// arbitrary inactive node to the nearest active node.
//
// Active nodes map to -1. Disconnected nodes don't appear in the map.
export type PrecalculatedAllocationPaths = Record<number, number>;

// Returns a path from targetNodeId to the nearest active node.
// Returns undefined if the target node is disconnected from the active nodes.
export function calculateAllocationPath(precalculatedPaths: PrecalculatedAllocationPaths, targetNodeId: number): number[] | undefined {
  const maxDepth = Object.keys(precalculatedPaths).length;
  if (maxDepth === 0) {
    return undefined;
  }

  const path = [targetNodeId];
  for (let i = 0; i <= maxDepth; i++) {
    const nextHop = precalculatedPaths[targetNodeId];
    if (nextHop === undefined) {
      // node is disconnected from activeNodes
      return undefined;
    }
    // assert((nextHop === -1) === activeNodes?.includes(nodeId))
    if (nextHop === -1) {
      return path;
    }
    path.push(nextHop);
    targetNodeId = nextHop;
  }

  throw new Error(`CalculateAllocationPaths produced a loop: ${JSON.stringify(path)}`);
}
