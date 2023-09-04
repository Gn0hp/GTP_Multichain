# @version ^0.3.8

OWNER: public(address)
TOKEN_NAME: public(String[50])
TOKEN_SYMBOL: public(String[50])
TOTAL_SUPPLY: public(uint256)
# DECIMALS: public(uint8)

_balances: HashMap[address, uint256]
_allowances: HashMap[address, HashMap[address, uint256]]
# event
event Transfer:
    _from: indexed(address)
    _to: indexed(address)
    _value: uint256
event Approval:
    _owner: indexed(address)
    _spender: indexed(address)
    _value: uint256
event ReceiveEther:
    _from: indexed(address)
    _amount: uint256
    _current_bal: uint256
    _gasLeft: uint256
event OwnershipTransferred:
    _oldOwner: indexed(address)
    _newOwner: indexed(address)

@external
def __init__(_name: String[50], _symbol: String[50], _totalSupply: uint256):
    self.OWNER = msg.sender
    self.TOKEN_NAME = _name
    self.TOKEN_SYMBOL = _symbol
    self.TOTAL_SUPPLY = 0

    self._mint(msg.sender, _totalSupply)

# ----------------- Ownable -----------------
@internal
def _transferOwnership(newOwner: address):
    oldOwner:address = self.OWNER
    self.OWNER = newOwner
    log OwnershipTransferred(oldOwner, newOwner)

@external
def transferOwnership(newOwner: address):
    assert newOwner != empty(address), "Ownable: Transfer ownership to address(0)"
    self._transferOwnership(newOwner)
# ----------------------- Token ERC20 -------------------------

@external
def balanceOf(account: address) -> uint256:
    return self._balances[account]
@external
def transfer(recipient: address, amount: uint256) -> bool:
    owner: address = msg.sender
    self._transfer(owner, recipient, amount)
    return True
@external
def allowance(owner: address, spender: address) -> uint256:
    return self._allowances[owner][spender]
@external
def approve(spender: address, amount: uint256) -> bool:
    owner: address = msg.sender
    self._approve(owner, spender, amount)
    return True
@external
def transferFrom(sender: address, recipient: address, amount: uint256) -> bool:
    spender: address = msg.sender
    self._spendAllowance(sender, spender, amount)
    self._transfer(sender, recipient, amount)
    return True
@external
def increaseAllowance(spender: address, addedValue: uint256) -> bool:
    self._approve(msg.sender, spender, self._allowances[msg.sender][spender] + addedValue)
    return True
@external
def decreaseAllowance(spender: address, subtractedValue: uint256) -> bool:
    currentAllowance: uint256 = self._allowances[msg.sender][spender]
    assert currentAllowance >= subtractedValue, "ERC20: decreased allowance below 0"
    self._approve(msg.sender, spender, currentAllowance - subtractedValue)
    return True

@internal
def _transfer(sender: address, recipient: address, amount: uint256):
    assert sender != empty(address), "ERC20: transfer from address(0)"
    assert recipient != empty(address), "ERC20; transfer to address(0)"

    self._beforeTokenTransfer(sender, recipient, amount)

    senderBalance: uint256 = self._balances[sender]
    assert senderBalance >= amount, "ERC20: transfer amount exceeds balance"
    self._balances[sender] -= amount
    self._balances[recipient] += amount
    log Transfer(sender, recipient, amount)

    self._afterTokenTransfer(sender, recipient, amount)

@internal
def _approve(owner: address, spender: address, amount: uint256):
    assert owner != empty(address), "ERC20: Invalid approver of address(0)"
    assert spender != empty(address), "ERC20: Invalid spender of address(0)"

    self._allowances[owner][spender] = amount
    log Approval(owner, spender, amount)
@internal
def _spendAllowance(owner: address, spender: address, amount: uint256):
    currentAllowance: uint256 = self._allowances[owner][spender]
    if currentAllowance != max_value(uint256):
        assert currentAllowance >= amount, "ERC20: Insufficient allowance"
        self._approve(owner, spender, currentAllowance - amount)

@internal
def _mint(account: address, amount: uint256):
    assert account != empty(address), "ERC20: mint to address(0)"

    self._beforeTokenTransfer(empty(address), account, amount)
    self._balances[account] += amount
    self.TOTAL_SUPPLY += amount

    log Transfer(empty(address), account, amount)
    self._afterTokenTransfer(empty(address), account, amount)

@internal
def _burn(account: address, amount: uint256):
    assert account != empty(address), "ERC20: burn from address(0)"
    assert self.TOTAL_SUPPLY >= amount, "ERC20: not enough balance"

    self._beforeTokenTransfer(account, empty(address), amount)
    self._balances[account] -= amount
    self.TOTAL_SUPPLY -= amount

    log Transfer(account, empty(address), amount)
    self._afterTokenTransfer(account, empty(address), amount)

@internal
def _beforeTokenTransfer(_from: address, _to: address, amount: uint256):
    pass

@internal
def _afterTokenTransfer(_from: address, _to: address, amount: uint256):
    pass

#-------------------- End of Token ERC20 ---------------------
@external
def mint(to: address, amount: uint256):
    assert self.OWNER == msg.sender,  "GToken: You are not allowed accessing this resource"
    self._mint(to, amount)
@external
def burn(account: address, amount: uint256):
    assert self.OWNER == msg.sender,  "GToken: You are not allowed accessing this resource"
    self._burn(account, amount)




# ----------------- Default functions---------------------
@external
@payable
# receive ether
def __default__():
    log ReceiveEther(msg.sender, msg.value, self.balance, msg.gas)

@external
# send ether to this contract then send to address
def sendEther(to: address, amount: uint256):
    send(to, amount)

@external
def sendAllEther(to: address):
    send(to, self.balance)

@external
def checkContractBalance() -> uint256:
    return self.balance