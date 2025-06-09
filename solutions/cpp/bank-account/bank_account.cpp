#include "bank_account.h"

namespace Bankaccount {

void Bankaccount::open( ) 
{
    if( account_open==true ) 
        throw std::runtime_error( "Attempted to open an already opened account." );
    account_open=true;
}

void Bankaccount::close( ) 
{
    if( account_open==true )
        throw std::runtime_error( "Attempted to close an already closed account." );
    account_open=false; 
    account_balance=0;
}

int  Bankaccount::balance( ) 
{ 
    if( account_open==false )
        throw std::runtime_error( "Attempted to check the balance of a closed account." );
    return account_balance; 
}

void Bankaccount::deposit(  int deposit_amount ) 
{
    account_mutex.lock();
    if( account_open==false )
    { // Attempted to depoit to a closed account
        account_mutex.unlock();
        throw std::runtime_error( "Attempted to depoit to a closed account." );
    } else if( deposit_amount<0  )
    { // Attempted to depoit a negative amount
        account_mutex.unlock();
        throw std::runtime_error( "Attempted to depoit a negative amount." );
    } 
    account_balance += deposit_amount;
    account_mutex.unlock();
}

void Bankaccount::withdraw( int withdrawal_amount ) 
{
    account_mutex.lock();
    if( account_open==false )
    { // Attempted to withdraw from a closed account
        account_mutex.unlock();
        throw std::runtime_error( "Attempted to withdraw from a closed account." );       
    } else if( withdrawal_amount>account_balance )
    { // Attempted to over-draw the account
        account_mutex.unlock();
        throw std::runtime_error( "Attempted to overdraw the account.");    
    } else if( withdrawal_amount<0  )
    { // Attempted to withdraw a negative amount
        account_mutex.unlock();
        throw std::runtime_error( "Attempted to withdraw a negative amount." );    
    } else
    {
        account_balance -= withdrawal_amount;
    }
    account_mutex.unlock();
}

} // namespace Bankaccount
