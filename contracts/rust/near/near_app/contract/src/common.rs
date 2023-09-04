use near_sdk::AccountId;
use near_sdk::borsh::{BorshDeserialize, BorshSerialize};

#[derive(Debug, BorshDeserialize, BorshSerialize)]
pub struct TransferEvent {
    pub from: AccountId,
    pub to: AccountId,
    pub amount: u128,
}
#[derive(Debug, BorshDeserialize, BorshSerialize)]
pub struct ApprovalEvent {
    pub owner: AccountId,
    pub spender: AccountId,
    pub amount: u128,
}
pub trait NEP20 {
    fn name(&self) -> String;

    fn symbol(&self) -> String;

    fn decimals(&self) -> u8;

    fn total_supply(&self) -> u128;

    fn balance_of(&self, account_id: AccountId) -> u128;

    fn transfer(&mut self, to: AccountId, amount: u128) -> bool;

    fn transfer_from(&mut self, from: AccountId, to: AccountId, amount: u128) -> bool;

    fn allowance(&self, owner: AccountId, spender: AccountId) -> u128;

    fn approve(&mut self, spender: AccountId, amount: u128) -> bool;
}

pub trait Ownable {
    fn owner(&self) -> AccountId;

    fn transfer_ownership(&mut self, new_owner: AccountId) -> bool;

    fn renounce_ownership(&mut self) -> bool;
}