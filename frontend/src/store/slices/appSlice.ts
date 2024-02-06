import { createSlice, PayloadAction } from '@reduxjs/toolkit'

export interface CounterState {
  value: number
}

const initialState: CounterState = {
  value: 0
}

export const counterSlice = createSlice({
    name: 'counter',
    initialState,
    reducers: {
      incrementByAmount: (state, action: PayloadAction<number>) => {
        state.value += action.payload
      }
    }
  })